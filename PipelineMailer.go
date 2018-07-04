package PluginOSS
import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/influxdata/kapacitor/tick/ast"
	"github.com/pkg/errors"

type AlertNode struct{ *AlertNodeData }

type AlertNodeData struct {

	Category string `json:"category"`
	Topic string `json:"topic"`
	}




func (n *AlertNodeData) Email(to ...string) *EmailHandler {
	em := &EmailHandler{
		AlertNodeData: n,
		ToList:        to,
	}
	n.EmailHandlers = append(n.EmailHandlers, em)
	return em
}

// Email AlertHandler
// tick:embedded:AlertNode.Email
type EmailHandler struct {
	*AlertNodeData `json:"-"`

	// List of email recipients.
	// tick:ignore
	ToList []string `tick:"To" json:"to"`
}