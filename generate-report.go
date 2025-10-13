package main

import (
    "errors"
    "fmt"
    "log"
    "os"
    "time"

    "github.com/joho/godotenv"

    "github.com/johnfercher/maroto/v2"
    
    // "github.com/johnfercher/maroto/v2/pkg/components/code"
    "github.com/johnfercher/maroto/v2/pkg/components/col"
    "github.com/johnfercher/maroto/v2/pkg/components/image"
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

func mkdirp(path string) {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}
}

func createReport() {
    t := time.Now()
    m := GetMaroto()
    timestamp := t.Format("2006-01")
    reportLocation := "./reports/caseload-" + timestamp + ".pdf"

    mkdirp("./reports")
    
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

    err := caseloadReport.RegisterHeader(getHeader())
    if err != nil {
        log.Fatal(err.Error())
    }

    t := time.Now()
    title := "Case Load Report: " + t.Format("Jan 2006")

    caseloadReport.AddRows(text.NewRow(10, title, props.Text{
        Size:  14,
        Top:   3,
        Style: fontstyle.Bold,
        Align: align.Center,
    }))

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
    contents := queryDb()

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

func getContents() [][]string {
    return [][]string{
        {"FirstName", "LastName"},
        {"John", "Doe"},
    }
}

func getHeader() core.Row {
    err := godotenv.Load()
    if err != nil {
	    log.Fatal("Error loading .env file")
    }


    business_address := os.Getenv("BUSINESS_ADDRESS")
    business_name    := os.Getenv("BUSINESS_NAME")
    business_number  := os.Getenv("BUSINESS_NUMBER")
    business_logo    := os.Getenv("BUSINESS_LOGO")
    business_website := os.Getenv("BUSINESS_WEBSITE")

    header := row.New(30).Add(
        image.NewFromFileCol(3, business_logo, props.Rect{
            Center:  true,
        }),
        col.New(6),
        col.New(3).Add(
            text.New(business_name, props.Text{
                Size:  8,
                Align: align.Right,
                Color: getRedColor(),
            }),
            text.New(business_address, props.Text{
                Top:   3,
                Size:  8,
                Align: align.Right,
                Color: getRedColor(),
            }),
            text.New(business_number, props.Text{
                Top:   12,
                Style: fontstyle.BoldItalic,
                Size:  8,
                Align: align.Right,
                Color: getBlueColor(),
            }),
            text.New(business_website, props.Text{
                Top:   15,
                Style: fontstyle.BoldItalic,
                Size:  8,
                Align: align.Right,
                Color: getBlueColor(),
            }),
        ),
    )
    
    return header
}

func getGrayColor() *props.Color {
    return &props.Color{
        Red:   200,
        Green: 200,
        Blue:  200,
    }
}

func getBlueColor() *props.Color {
    return &props.Color{
        Red:   10,
        Green: 10,
        Blue:  150,
    }
}

func getRedColor() *props.Color {
    return &props.Color{
        Red:   150,
        Green: 10,
        Blue:  10,
    }
}

