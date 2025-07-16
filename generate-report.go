package main

import (
    "fmt"
    "log"
    "time"

    "github.com/johnfercher/maroto/v2"
    "github.com/johnfercher/maroto/v2/pkg/components/code"
    "github.com/johnfercher/maroto/v2/pkg/config"
    "github.com/johnfercher/maroto/v2/pkg/consts/orientation"
    "github.com/johnfercher/maroto/v2/pkg/consts/pagesize"
    "github.com/johnfercher/maroto/v2/pkg/core"
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

    mrt := maroto.New(cfg)

    mrt.AddRow(20,
            code.NewBarCol(4, "barcode"),
            code.NewMatrixCol(4, "matrixcode"),
            code.NewQrCol(4, "qrcode"),
        )

    return mrt
}
