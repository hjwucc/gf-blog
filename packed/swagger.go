package packed

import "github.com/gogf/gf/os/gres"

func init() {
	if err := gres.Add("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"); err != nil {
		panic("add binary content to resource manager failed: " + err.Error())
	}
}
