package node

import (
	"errors"
)

// Dimension struct
type Dimension struct {
	Key   string `json:"key"`
	Value string `json:"val"`
}

// Metric struct
type Metric struct {
	Key   string `json:"key"`
	Value int    `json:"val"`
}

// WebTraffic struct
type WebTraffic struct {
	Dimensions []Dimension `json:"dim"`
	Metrics    []Metric    `json:"metrics"`
}

//Response holds api response
type Response struct {
	Errors  []string `json:"errors,omitempty"`
	Message string   `json:"message,omitempty"`
}

// Node struct
type Node struct {
	content  map[string]interface{}
	parent   *Node
	children []*Node
}

// New creates a new node
func New(parent *Node) *Node {
	return &Node{
		content: make(map[string]interface{}),
		parent:  parent,
	}
}

// IsRootNode returns true if parent node
func (n *Node) IsRootNode() bool {
	return n.parent == nil
}

// NewChild creates a new child node with the current node as parent
func (n *Node) NewChild() *Node {
	child := New(n)
	n.children = append(n.children, child)
	return child
}

// GetChildren retrieves all the children
func (n *Node) GetChildren() []*Node {
	return n.children
}

// GetDimension gets the dimension value of specified key
func (n *Node) GetDimension(key string) string {
	value, _ := n.content[key]
	if value == nil {
		return ""
	}
	str, _ := value.(string)
	return str
}

// SetContent sets the content with the specified key and its value
func (n *Node) SetContent(key string, value interface{}) {
	n.content[key] = value
}

// GetMetric gets the values of specified metric key
func (n *Node) GetMetric(key string) int {
	value, _ := n.content[key]
	if value == nil {
		return 0
	}
	val, _ := value.(int)
	return val
}

// SetMetric the content key and its value
func (n *Node) SetMetric(key string, value int) {
	p := n
	for p != nil {
		existingValue := p.GetMetric(key)
		p.SetContent(key, value+existingValue)
		p = p.parent
	}
}

// // GetParent retrieves the parent node
// func (n *Node) GetParent() *Node {
// 	return n.parent
// }

// IsCountryNode returns true if parent node
// func (n *Node) IsCountryNode() bool {
// 	_, ok := n.content["country"]
// 	return ok
// }

// IsDeviceNode returns true if parent node
// func (n *Node) IsDeviceNode() bool {
// 	_, ok := n.content["device"]
// 	return ok
// }

// UpdateMetric func
func (n *Node) UpdateMetric(country string, device string, metric string, value int) error {
	if !n.IsRootNode() {
		return errors.New("This operation is allowed only on root node")
	}
	var (
		countryNode *Node
		deviceNode  *Node
	)
	//looping through country level children to find the node
	for _, c := range n.GetChildren() {
		if c.GetDimension("country") == country {
			countryNode = c
		}
	}
	//Create new child node in country level if not present
	if countryNode == nil {
		countryNode = n.NewChild()
		countryNode.SetContent("country", country)
	}
	//looping through device level children to find the node
	for _, d := range countryNode.GetChildren() {
		if d.GetDimension("device") == device {
			deviceNode = d
		}
	}
	//Create new child node in device level if not present
	if deviceNode == nil {
		deviceNode = countryNode.NewChild()
		deviceNode.SetContent("device", device)
	}
	//Set the metric value in the device node
	deviceNode.SetMetric(metric, value)
	return nil
}

//GetMetricByCountry gets metrics by country
func (n *Node) GetMetricByCountry(request WebTraffic, country string) (WebTraffic, error) {
	if !n.IsRootNode() {
		return WebTraffic{}, errors.New("This operation is allowed only on root node")
	}
	var (
		countryNode    *Node
		countryTraffic WebTraffic
	)
	countryTraffic.Dimensions = request.Dimensions
	countryTraffic.Metrics = []Metric{}
	//Get the node of the specified country
	for _, c := range n.GetChildren() {
		if c.GetDimension("country") == country {
			countryNode = c
			break
		}
	}
	if countryNode == nil {
		return countryTraffic, nil
	}
	//Append the metrics values
	countryTraffic.Metrics = append(countryTraffic.Metrics, Metric{Key: "webreq", Value: countryNode.GetMetric("webreq")}, Metric{Key: "timespent", Value: countryNode.GetMetric("timespent")})
	return countryTraffic, nil
}
