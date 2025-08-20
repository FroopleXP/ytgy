package ytgy

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
    hms
    p10
    vts
    vts1
    vts2
    ms
    ms1
    ms2
    msv
    msv1
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
    case hms:
        return "%02d%02d%02d"
    case p10:
        return "P10%05d"
    case vts:
        return "VTS %02d %01d"
    case vts1:
        return "VTS %03d 1"
    case vts2:
        return "VTS 01 %03d"
    case ms:
        return "My Slideshow"
    case ms1:
        return "My Slideshow %02d"
    case ms2:
        return "My Slideshow Video"
    case msv:
        return "My Stupeflix Video"
    case msv1:
        return "My Stupeflix Video %04d"
    }
    return ""
}

func (p pattern) Source() string {
    switch p {
    case mov, onehun, sam, dsc, sdv, dscf, dscn, pict, maq, mol, p10:
        return "Camera"
    case file:
        return "Dashcam"
    case gopr, gp, gx:
        return "GoPro"
    case dji:
        return "Drone"
    case hni:
        return "Nintendo DS"
    case wa, hms:
        return "Misc."
    case ms, ms1, ms2, msv, msv1:
        return "Video Editor"
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
    case hms:
        return fmt.Sprintf(p.String(), rand.Intn(24), rand.Intn(60), rand.Intn(60))
    case p10:
        return fmt.Sprintf(p.String(), rand.Intn(20_000))
    case vts:
        return fmt.Sprintf(p.String(), rand.Intn(100), rand.Intn(10))
    case vts1, vts2:
        return fmt.Sprintf(p.String(), rand.Intn(1_000))
    case ms1:
        return fmt.Sprintf(p.String(), rand.Intn(100))
    case msv1:
        return fmt.Sprintf(p.String(), rand.Intn(1051))
    }
    return p.String()
}

func (p pattern) Url() string {
    return fmt.Sprintf(ytBaseUrl, url.PathEscape(p.Generate()))
}

func Rand() pattern {
    return pattern(rand.Intn(int(total)))
}

