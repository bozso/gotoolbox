package table

type (
    Row []string
    Header []string
)

type Table struct{
    Header     `json:"header"`
    Rows []Row `json:"rows"`
}

func (t Table) WriteTo(w Writer) {
    w.AddHeader(t.Header)
    
    for _, row := range t.Rows {
        w.AddRow(row)
    }
}

func (t Table) ToString(mode string) (s string, err error) {
    var tmode Mode
    if err = tmode.Set(mode); err != nil {
        return
    }
    
    var w Writer
    switch tmode {
    case Html:
        w = &HtmlWriter{}
    case Latex:
        w = &LatexWriter{}   
    }
    
    t.WriteTo(w)
    
    s = w.String()
    return
}
