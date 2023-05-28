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
	Language string
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
			o.Run(o.Number, o.Keywords, o.Language)
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
	cmd.Flags().StringVarP(&o.Language, "language", "l", "日本語", "Language")
	cmd.MarkFlagRequired("keywords")

	return cmd
}

func (o *GentitleOptions) Run(n int, kw []string, l string) {
	sysMsg := `
あなたはアウトプットが得意なソフトウェアエンジニアです。
他のエンジニアがテックブログを書こうとしているので質問に対して適切なアドバイスを返してください。
`
	usrMsg := fmt.Sprintf(
		`
下記のキーワードに関連するテックブログに適した、簡潔で信頼性があり注目を集めるタイトルを考えてください。

%s

タイトルの例を%d個挙げてください。
言語は%sでお願いします。
`,
		strings.Join(kw, "\n"),
		n,
		l,
	)

	resp := o.ChatGPTClient.Chat(sysMsg, usrMsg)
	fmt.Println(resp)
}
