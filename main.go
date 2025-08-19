package main

import (
    "fmt"
    "net/url"
    "math/rand"
)

const ytBaseUrl = "https://www.youtube.com/results?search_query=%s"

type pattern int

const (
    img = pattern(iota)
    mvi
    mov
    onehun
    sam
    dsc
    sdv
    dscf
    dscn
    pict
    maq
    file
    gopr
    gp
    gx
    dji
    hni
    wa
    mol
    // NOTE: Place all patterns above this line!
    total
)

func (p pattern) String() string {
    switch p {
    case img:
        return "IMG %04d"
    case mvi:
        return "MVI %04d"
    case mov:
        return "MOV %04d"
    case onehun:
        return "100 %04d"
    case sam:
        return "SAM %04d"
    case dsc:
        return "DSC %04d"
    case sdv:
        return "SDV %04d"
    case dscf:
        return "DSCF%04d"
    case dscn:
        return "DSCN%04d"
    case pict:
        return "PICT%04d"
    case maq:
        return "MAQ0%04d"
    case file:
        return "FILE%04d"
    case gopr:
        return "GOPR%04d"
    case gp:
        return "GP01%04d"
    case gx:
        return "GX01%04d"
    case dji:
        return "DJI %04d"
    case hni:
        return "HNI 0%03d"
    case wa:
        return "WA0%03d" 
    case mol:
        return "MOL0%02X"
    }
    return ""
}

func (p pattern) Source() string {
    switch p {
    case mov, onehun, sam, dsc, sdv, dscf, dscn, pict, maq, mol:
        return "Camera"
    case file:
        return "Dashcam"
    case gopr, gp, gx:
        return "GoPro"
    case dji:
        return "Drone"
    case hni:
        return "Nintendo DS"
    case wa:
        return "Misc."
    }
    return "Unknown"
}

func (p pattern) Generate() string {
    switch p {
    case img, mvi, mov, onehun, sam, dsc, sdv, dscf, dscn, pict, maq, file, gopr, gp, gx, dji:
        return fmt.Sprintf(p.String(), rand.Intn(10_000))
    case hni:
        return fmt.Sprintf(p.String(), rand.Intn(101))
    case wa:
        return fmt.Sprintf(p.String(), rand.Intn(1_000))
    case mol:
        return fmt.Sprintf(p.String(), rand.Intn(100))
    }
    return ""
}

func (p pattern) Url() string {
    return fmt.Sprintf(ytBaseUrl, url.PathEscape(p.Generate()))
}

func Rand() pattern {
    return pattern(rand.Intn(int(total)))
}

func main() {
    r := Rand()
    fmt.Printf("Go search '%s', device is usually %s (%s)\n", r.Generate(), r.Source(), r.Url())
}
