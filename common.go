package main

import (
	"encoding/hex"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strconv"
)

func PrintHeader() error {
	assets, err := NewAssetsReader(os.Args[2])
	if err != nil {
		return err
	}
	defer assets.Close()

	log.Print(dump(assets.Header))

	return nil
}

func PrintMeta() error {
	assets, err := NewAssetsReader(os.Args[2])
	if err != nil {
		return err
	}
	defer assets.Close()

	var filterByType uint64
	if len(os.Args) > 3 {
		filterByType, err = strconv.ParseUint(os.Args[3], 10, 32)
		return errors.Errorf("invalid type_id: %v", filterByType)
	}

	return assets.RangeObjects(func(desc Object, r io.Reader) error {
		if filterByType != 0 && desc.TypeID != uint32(filterByType) {
			return nil
		}

		log.Printf("%+v", desc)
		return nil
	})
}

func PrintHexDump() error {
	assets, err := NewAssetsReader(os.Args[2])
	if err != nil {
		return err
	}
	defer assets.Close()

	var filterByType uint64
	if len(os.Args) > 3 {
		filterByType, err = strconv.ParseUint(os.Args[3], 10, 32)
		return errors.Errorf("invalid type_id: %v", filterByType)
	}

	return assets.RangeObjects(func(desc Object, r io.Reader) error {
		if filterByType != 0 && desc.TypeID != uint32(filterByType) {
			return nil
		}
		data, err := ioutil.ReadAll(r)
		if err != nil {
			return err
		}
		log.Printf("%+v\n%v", desc, hex.Dump(data))
		return nil
	})
}

func UnpackAssets() error {
	assets, err := NewAssetsReader(os.Args[2])
	if err != nil {
		return err
	}
	defer assets.Close()

	err = os.MkdirAll(os.Args[3], 0777)
	if err != nil {
		return err
	}

	return assets.RangeObjects(func(desc Object, r io.Reader) error {
		dir := path.Join(os.Args[3], strconv.FormatUint(uint64(desc.TypeID), 10))
		err := os.MkdirAll(dir, 0777)
		if err != nil {
			return err
		}

		file := path.Join(dir, strconv.FormatUint(uint64(desc.ID), 10))
		out, err := os.Create(file)
		if err != nil {
			return err
		}

		_, err = io.Copy(out, r)
		if err != nil {
			out.Close()
			return err
		}

		return out.Close()
	})
}