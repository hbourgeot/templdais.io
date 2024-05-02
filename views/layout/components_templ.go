// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.663
package layout

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"
import "strings"

import "github.com/hbourgeot/templdais"

var menu = templdais.MenuAttrs{
	Items: []templdais.MenuItem{
		{Text: "string", Link: templ.SafeURL("/"), Class: "w-full"},
		{Text: "string", Link: templ.SafeURL("/"), Class: "w-full"},
		{Text: "string", Link: templ.SafeURL("/"), Class: "w-full"},
		{Text: "string", Link: templ.SafeURL("/"), Class: "w-full"},
		{Text: "string", Link: templ.SafeURL("/"), Class: "w-full"},
		{Text: "string", Link: templ.SafeURL("/"), Class: "w-full"},
		{Text: "string", Link: templ.SafeURL("/"), Class: "w-full"},
		{Text: "string", Link: templ.SafeURL("/"), Class: "w-full"},
		{Text: "string", Link: templ.SafeURL("/"), Class: "w-full"},
		{Text: "string", Link: templ.SafeURL("/"), Class: "w-full"},
		{Text: "string", Link: templ.SafeURL("/"), Class: "w-full"},
		{Text: "string", Link: templ.SafeURL("/"), Class: "w-full"},
		{Text: "string", Link: templ.SafeURL("/"), Class: "w-full"},
		{Text: "string", Link: templ.SafeURL("/"), Class: "w-full"},
		{Text: "string", Link: templ.SafeURL("/"), Class: "w-full"},
		{Text: "string", Link: templ.SafeURL("/"), Class: "w-full"},
		{Text: "string", Link: templ.SafeURL("/"), Class: "w-full"},
		{Text: "string", Link: templ.SafeURL("/"), Class: "w-full"},
		{Text: "string", Link: templ.SafeURL("/"), Class: "w-full"},
		{Text: "string", Link: templ.SafeURL("/"), Class: "w-full"},
		{Text: "string", Link: templ.SafeURL("/"), Class: "w-full"},
		{Text: "string", Link: templ.SafeURL("/"), Class: "w-full"},
		{Text: "string", Link: templ.SafeURL("/"), Class: "w-full"},
		{Text: "string", Link: templ.SafeURL("/"), Class: "w-full"},
		{Text: "string", Link: templ.SafeURL("/"), Class: "w-full"},
		{Text: "string", Link: templ.SafeURL("/"), Class: "w-full"},
		{Text: "string", Link: templ.SafeURL("/"), Class: "w-full"},
		{Text: "string", Link: templ.SafeURL("/"), Class: "w-full"},
		{Text: "string", Link: templ.SafeURL("/"), Class: "w-full"},
		{Text: "string", Link: templ.SafeURL("/"), Class: "w-full"},
		{Text: "string", Link: templ.SafeURL("/"), Class: "w-full"},
		{Text: "string", Link: templ.SafeURL("/"), Class: "w-full"},
		{Text: "string", Link: templ.SafeURL("/"), Class: "w-full"},
		{Text: "string", Link: templ.SafeURL("/"), Class: "w-full"},
		{Text: "string", Link: templ.SafeURL("/"), Class: "w-full"},
		{Text: "string", Link: templ.SafeURL("/"), Class: "w-full"},
		{Text: "string", Link: templ.SafeURL("/"), Class: "w-full"},
		{Text: "string final", Link: templ.SafeURL("/"), Class: "w-full"},
	},
	Vertical:   true,
	Size:       "lg",
	Responsive: false,
	Class:      "bg-base-600 w-1/5 fixed top-0 left-0 w-64 h-screen overflow-x-hidden !overflow-y-auto !flex !flex-col !flex-nowrap",
}

func Section() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`min-height:calc(100vh - 71px);`)
	templ_7745c5c3_CSSID := templ.CSSID(`Section`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func Components(components ...templ.Component) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		var templ_7745c5c3_Var2 = []any{"flex justify-start gap-4 w-full relative", Section()}
		templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var2...)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<section class=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 string
		templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var2).String())
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/layout/components.templ`, Line: 1, Col: 0}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templdais.Menu(menu).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for _, c := range components {
			templ_7745c5c3_Err = c.Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</section>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
