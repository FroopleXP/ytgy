package main

import (
    "fmt"
    "math/rand"
)


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

func Rand() pattern {
    return pattern(rand.Intn(int(total)))
}

func main() {
    fmt.Println(mol.Generate())
}
