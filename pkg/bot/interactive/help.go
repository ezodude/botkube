package interactive

import (
	"fmt"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/kubeshop/botkube/pkg/config"
)

// RunCommandName defines the button name for the run commands.
const RunCommandName = "Run command"

// HelpMessage provides an option to build the Help message depending on a given platform.
type HelpMessage struct {
	btnBuilder  ButtonBuilder
	botName     string
	platform    config.CommPlatformIntegration
	clusterName string
}

// NewHelpMessage return a new instance of HelpMessage.
func NewHelpMessage(platform config.CommPlatformIntegration, clusterName, botName string) *HelpMessage {
	btnBuilder := ButtonBuilder{BotName: botName}
	return &HelpMessage{btnBuilder: btnBuilder, botName: botName, platform: platform, clusterName: clusterName}
}

// Build returns help message with interactive sections.
func (h *HelpMessage) Build() Message {
	msg := Message{
		Base: Base{
			Description: fmt.Sprintf("Botkube is now active for %q cluster :rocket:", h.clusterName),
		},
	}

	type getter func() []Section
	var sections = []getter{
		h.cluster,
		h.notificationSections,
		h.kubectlSections,
		h.filters,
		h.feedback,
		h.footer,
	}
	for _, add := range sections {
		msg.Sections = append(msg.Sections, add()...)
	}

	return msg
}

func (h *HelpMessage) cluster() []Section {
	return []Section{
		{
			Base: Base{
				Header:      "Using multiple instances",
				Description: fmt.Sprintf("If you are running multiple Botkube instances in the same channel to interact with %s, make sure to specify the cluster name when typing commands.", h.clusterName),
				Body: Body{
					CodeBlock: fmt.Sprintf("--cluster-name=%s\n", h.clusterName),
				},
			},
		},
		{
			Base: Base{
				Header:      "Ping your cluster",
				Description: "Check the status of connected Kubernetes cluster(s).",
			},
			Buttons: []Button{
				h.btnBuilder.ForCommandWithDescCmd("Check status", "ping"),
			},
		},
	}
}

func (h *HelpMessage) filters() []Section {
	return []Section{
		{
			Base: Base{
				Header: "Filters (advanced)",
				Body: Body{
					Plaintext: "You can extend Botkube functionality by writing additional filters that can check resource specs, validate some checks and add messages to the Event struct. Learn more at https://botkube.io/filters",
				},
			},
		},
	}
}

func (h *HelpMessage) feedback() []Section {
	return []Section{
		{
			Base: Base{
				Header: "Angry? Amazed?",
			},
			Buttons: []Button{
				h.btnBuilder.DescriptionURL("Give feedback", "feedback", "https://feedback.botkube.io", ButtonStylePrimary),
			},
		},
	}
}

func (h *HelpMessage) footer() []Section {
	return []Section{
		{
			Buttons: []Button{
				h.btnBuilder.ForURL("Read our docs", "https://botkube.io/docs"),
				h.btnBuilder.ForURL("Join our Slack", "https://join.botkube.io"),
				h.btnBuilder.ForURL("Follow us on Twitter", "https://twitter.com/botkube_io"),
			},
		},
	}
}

func (h *HelpMessage) notificationSections() []Section {
	return []Section{
		{
			Base: Base{
				Header: "Manage incoming notifications",
				Body: Body{
					CodeBlock: fmt.Sprintf("%s notifier [start|stop|status]\n", h.botName),
				},
			},
			Buttons: []Button{
				h.btnBuilder.ForCommandWithoutDesc("Start notifications", "notifier start"),
				h.btnBuilder.ForCommandWithoutDesc("Stop notifications", "notifier stop"),
				h.btnBuilder.ForCommandWithoutDesc("Get status", "notifier status"),
			},
		},
		{
			Base: Base{
				Header:      "Notification settings for this channel",
				Description: "By default, Botkube will notify only about cluster errors and recommendations.",
			},
			Buttons: []Button{
				h.btnBuilder.ForCommandWithDescCmd("Adjust notifications", "edit SourceBindings", ButtonStylePrimary),
			},
		},
	}
}

func (h *HelpMessage) kubectlSections() []Section {
	if h.platform == config.SocketSlackCommPlatformIntegration {
		return []Section{
			{
				Base: Base{
					Header: "Interactive kubectl - no typing!",
				},
				Buttons: []Button{
					h.btnBuilder.ForCommandWithDescCmd("kubectl", "kubectl", ButtonStylePrimary),
				},
			},
			{
				Base: Base{
					Description: "Alternatively use kubectl as usual with all supported commands",
				},
				Buttons: []Button{
					h.btnBuilder.ForCommand("List commands", "commands list", "k | kc | kubectl [command] [options] [flags]"),
				},
			},
		}
	}

	// without the kubectl command builder
	return []Section{
		{
			Base: Base{
				Header:      "Run kubectl commands (if enabled)",
				Description: fmt.Sprintf("You can run kubectl commands directly from %s!", cases.Title(language.English).String(string(h.platform))),
			},
			Buttons: []Button{
				h.btnBuilder.ForCommandWithDescCmd(RunCommandName, "kubectl get services"),
				h.btnBuilder.ForCommandWithDescCmd(RunCommandName, "kubectl get pods"),
				h.btnBuilder.ForCommandWithDescCmd(RunCommandName, "kubectl get deployments"),
			},
		},
		{
			Base: Base{
				Description: "To list all supported kubectl commands",
			},
			Buttons: []Button{
				h.btnBuilder.ForCommandWithDescCmd("List commands", "commands list"),
			},
		},
	}
}
