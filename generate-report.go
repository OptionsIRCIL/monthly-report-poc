package main

import (
    "fmt"
    "log"
    "time"

    "github.com/johnfercher/maroto/v2"
    
    // "github.com/johnfercher/maroto/v2/pkg/components/code"
    "github.com/johnfercher/maroto/v2/pkg/components/col"
    // "github.com/johnfercher/maroto/v2/pkg/components/image"
    "github.com/johnfercher/maroto/v2/pkg/components/row"
    "github.com/johnfercher/maroto/v2/pkg/components/text"
    "github.com/johnfercher/maroto/v2/pkg/consts/align"
    "github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
    "github.com/johnfercher/maroto/v2/pkg/consts/orientation"
    "github.com/johnfercher/maroto/v2/pkg/consts/pagesize"

    "github.com/johnfercher/maroto/v2/pkg/config"
    "github.com/johnfercher/maroto/v2/pkg/core"
    "github.com/johnfercher/maroto/v2/pkg/props"
)

func createReport() {
    t := time.Now()
    m := GetMaroto()
    timestamp := t.Format("2006-01")
    reportLocation := "reports/caseload-" + timestamp + ".pdf"

    document, err := m.Generate()
    if err != nil {
        log.Fatal(err.Error())
    }

    err = document.Save(reportLocation)
    if err != nil {
        log.Fatal(err.Error())
    }

    fmt.Println("Report '" + reportLocation + "' created!")
}

func GetMaroto() core.Maroto {
    cfg := config.NewBuilder().
        WithPageNumber().
        WithOrientation(orientation.Vertical).
        WithPageSize(pagesize.Letter).
        WithLeftMargin(15).
        WithTopMargin(15).
        WithRightMargin(15).
        WithBottomMargin(15).
        Build()

    caseloadReport := maroto.New(cfg)

    caseloadReport.AddRows(buildTable()...)

    return caseloadReport
}

func buildTable() []core.Row {
    const COLUMNS = 2

    rows := []core.Row{
        row.New(5).Add(
            col.New(COLUMNS),
            text.NewCol(4, "First Name", props.Text{Size: 9, Align: align.Center, Style: fontstyle.Bold}),
            text.NewCol(2, "Last Name", props.Text{Size: 9, Align: align.Center, Style: fontstyle.Bold}),
        ),
    }

    var contentsRow []core.Row
    contents := getContents()

    for i, content := range contents {
        r := row.New(4).Add(
            col.New(COLUMNS),
            text.NewCol(4, content[0], props.Text{Size: 8, Align: align.Center}),
            text.NewCol(2, content[1], props.Text{Size: 8, Align: align.Center}),
        )
        if i % 2 == 0 {
            gray := getGrayColor()
            r.WithStyle(&props.Cell{BackgroundColor: gray})
        }

        contentsRow = append(contentsRow, r)
    }

    rows = append(rows, contentsRow...)

    return rows
}

func getGrayColor() *props.Color {
    return &props.Color{
        Red:   200,
        Green: 200,
        Blue:  200,
    }
}


func getContents() [][]string {
    return [][]string{
        {"FirstName", "LastName"},
        {"John", "Doe"},
    }
}

type Employee struct {
    FirstName string
    LastName string
}

