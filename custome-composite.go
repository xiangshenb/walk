package walk

import (
    "github.com/lxn/win"
)

const customCompositeWindowClass = `\o/ Walk_Custom_Composite_Class \o/`

func init() {
    MustRegisterWindowClass(customCompositeWindowClass)
}

type CustomComposite struct {
    ContainerBase
    brush Brush
}

func (cc *CustomComposite)WndProc(hwnd win.HWND, msg uint32, wParam, lParam uintptr) uintptr {
    switch msg {
    case win.WM_CTLCOLORSTATIC:
        children := cc.Children()
        l := children.Len()
        for i:=0;i<l;i++ {
            item := children.At(i)
            switch lab := item.(type) {
            case *Label:
                if item.Handle() == win.HWND(lParam) {
                    hdcStatic := win.HDC(wParam)
                    win.SetTextColor( hdcStatic, win.COLORREF(lab.TextColor) )
                    win.SetBkMode(hdcStatic, win.TRANSPARENT)
                }
            }
        }
        return uintptr(cc.brush.handle())
    case win.WM_DESTROY:
        cc.brush.Dispose()
    }
    return cc.ContainerBase.WndProc(hwnd, msg, wParam, lParam)
}

func newCustomCompositeWithStyle(parent Window, style uint32) (*CustomComposite, error) {
    c := new(CustomComposite)
    c.brush,_ = NewSolidColorBrush(RGB(0,0,0))
    c.children = newWidgetList(c)
    c.SetPersistent(true)

    if err := InitWidget(
        c,
        parent,
        customCompositeWindowClass,
        win.WS_CHILD|win.WS_VISIBLE|style,
        win.WS_EX_CONTROLPARENT); err != nil {
        return nil, err
    }

    return c, nil
}

func NewCustomComposite(parent Container) (*CustomComposite, error) {
    return newCustomCompositeWithStyle(parent, 0)
}
