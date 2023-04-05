package grafana

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

var errNotFound = errors.New("resource not found")

type Client struct {
	cli     *http.Client
	baseUrl string
	token   string
}

func NewClient(baseUrl string, token string) *Client {
	return &Client{
		cli:     http.DefaultClient,
		baseUrl: baseUrl,
		token:   token,
	}
}

func (c *Client) SaveDashboard(d *Dashboard) error {
	req := SaveDashboardRequest{
		Dashboard: d,
		Overwrite: true,
	}
	resp := &SaveDashboardResponse{}
	url := c.baseUrl + "/api/dashboards/db"
	return c.doWithBody("POST", url, req, resp)
}

func (c *Client) GetLibraryPanelByUID(uid string) (*PanelLibraryResponse, error) {
	resp := &PanelLibraryResponse{}
	url := c.baseUrl + "/api/library-elements/" + uid
	err := c.doNoBody("GET", url, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) CreateLibraryPanel(panel *Panel, folderUID string) (*PanelLibraryResponse, error) {
	//panel.LibraryPanel = LibraryPanel{
	//	Uid:  panel.UID,
	//	Name: panel.Title,
	//}
	req := &SaveLibraryPanelRequest{
		Name:      panel.Title,
		Model:     panel,
		FolderUid: folderUID,
		Kind:      1,
		Uid:       panel.UID,
	}
	resp := &PanelLibraryResponse{}
	url := c.baseUrl + "/api/library-elements"
	err := c.doWithBody("POST", url, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) UpdateLibraryPanel(panel *Panel, id int, version int, folderUID string) (*PanelLibraryResponse, error) {
	//panel.LibraryPanel = LibraryPanel{
	//	Uid:  panel.UID,
	//	Name: panel.Title,
	//}
	req := &SaveLibraryPanelRequest{
		Name:      panel.Title,
		Model:     panel,
		FolderUid: folderUID,
		Kind:      1,
		Uid:       panel.UID,
		Version:   version,
	}
	resp := &PanelLibraryResponse{}
	url := c.baseUrl + "/api/library-elements/" + panel.UID
	err := c.doWithBody("PATCH", url, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) DeleteLibraryPanel(uid string) (int64, error) {
	resp := make(map[string]any)
	url := c.baseUrl + "/api/library-elements/" + uid
	err := c.doNoBody("DELETE", url, &resp)
	if err != nil {
		return 0, err
	}
	return int64(resp["id"].(float64)), nil
}

func (c *Client) SaveLibraryPanel(panel *Panel, folderUID string) (*PanelLibraryResponse, error) {
	plResp, err := c.GetLibraryPanelByUID(panel.UID)
	if err != nil && err != errNotFound {
		return nil, err
	}
	if err == errNotFound {
		return c.CreateLibraryPanel(panel, folderUID)
	}
	return c.UpdateLibraryPanel(panel, plResp.Result.Id, plResp.Result.Version, folderUID)
}

func (c *Client) doNoBody(method, u string, response any) error {
	r, err := http.NewRequest(method, u, nil)
	if err != nil {
		return err
	}
	c.setHeaders(r)
	resp, err := c.cli.Do(r)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode == 404 {
		return errNotFound
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("save error: %v", string(body))
	}
	return json.Unmarshal(body, response)
}

func (c *Client) doWithBody(method, u string, request, response any) error {
	data, err := json.MarshalIndent(request, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(data))
	r, err := http.NewRequest(method, u, bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	c.setHeaders(r)
	resp, err := c.cli.Do(r)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("save error: %v", string(body))
	}
	fmt.Println(string(body))
	return json.Unmarshal(body, response)
}

func (c *Client) setHeaders(r *http.Request) {
	r.Header.Set("Authorization", "Bearer "+c.token)
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Accept", "application/json")
}
