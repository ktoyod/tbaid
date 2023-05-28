package gentitle

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/ktoyod/tbaid/pkg/chatgpt"
)

type GentitleOptions struct {
	ChatGPTClient *chatgpt.ChatGPTClient

	Number   int
	Keywords []string
}

func NewGentitleOptions() *GentitleOptions {
	c := chatgpt.NewChatGPTClient()

	return &GentitleOptions{
		ChatGPTClient: c,
	}
}

func NewCmdGentitle() *cobra.Command {
	o := NewGentitleOptions()

	cmd := &cobra.Command{
		Use:   "gentitle",
		Short: "Generate titles for tech blog",
		Run: func(cmd *cobra.Command, args []string) {
			o.Run(o.Number, o.Keywords)
		},
	}

	cmd.Flags().IntVarP(&o.Number, "number", "n", 3, "Number of titles to generate")
	cmd.Flags().StringSliceVarP(
		&o.Keywords,
		"keywords",
		"k",
		[]string{},
		"The keywords of the tech blog content. This flag takes comma-separated value as arguments and split them accordingly.",
	)

	return cmd
}

func (o *GentitleOptions) Run(n int, kw []string) {
	sysMsg := `
あなたはアウトプットが得意なソフトウェアエンジニアです。
他のエンジニアがテックブログを書こうとしているので質問に対して適切なアドバイスを返してください。
`
	usrMsg := fmt.Sprintf(
		`
テックブログを書いているのですが、簡潔で信頼性があり注目を集めるタイトルを考えてください。関連するキーワードは以下になります。

%s

タイトルの例を%d個挙げてください。
`,
		strings.Join(kw, "\n"),
		n,
	)

	resp := o.ChatGPTClient.Chat(sysMsg, usrMsg)
	fmt.Println(resp)
}
