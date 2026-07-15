package main

import (
	"context"
	"encoding/json"
	"log"
	"math"
	"net/http"

	"github.com/dianich2/jsonrpc2"
	"github.com/gorilla/mux"
)

type ParmsXY struct {
	X *float64 `json:"x"`
	Y *float64 `json:"y"`
}

type ParmN struct {
	N *float64 `json:"N"`
}

var serv = jsonrpc2.New()
var precision = 2

func roundToPrecision(val float64, num int) float64 {
	pow := math.Pow(10, float64(num))
	return math.Round(val*pow) / pow
}

func regMethods() {
	err := serv.Register("sum", func(ctx context.Context, params json.RawMessage) (interface{}, *jsonrpc2.Error) {
		nums, err := jsonrpc2.DecodeParams[[]float64](params)
		if err == nil {
			if len(nums) != 2 {
				return nil, jsonrpc2.ErrInvalidParams("you must input 2 params")
			}
			return roundToPrecision(nums[0]+nums[1], precision), nil
		}
		nums2, err := jsonrpc2.DecodeParams[ParmsXY](params)
		if nums2.X == nil || nums2.Y == nil {
			return nil, jsonrpc2.ErrInvalidParams("params object must contain x and y")
		}
		if err == nil {
			return roundToPrecision(*nums2.X+*nums2.Y, precision), nil
		}
		return nil, jsonrpc2.ErrInvalidParams("incorrect params")
	})
	if err != nil {
		log.Fatal(err)
	}

	err = serv.Register("sub", func(ctx context.Context, params json.RawMessage) (interface{}, *jsonrpc2.Error) {
		nums, err := jsonrpc2.DecodeParams[[]float64](params)
		if err == nil {
			if len(nums) != 2 {
				return nil, jsonrpc2.ErrInvalidParams("you must input 2 params")
			}
			return roundToPrecision(nums[0]-nums[1], precision), nil
		}
		nums2, err := jsonrpc2.DecodeParams[ParmsXY](params)
		if nums2.X == nil || nums2.Y == nil {
			return nil, jsonrpc2.ErrInvalidParams("params object must contain x and y")
		}
		if err == nil {
			return roundToPrecision(*nums2.X-*nums2.Y, precision), nil
		}
		return nil, jsonrpc2.ErrInvalidParams("incorrect params")
	})
	if err != nil {
		log.Fatal(err)
	}

	err = serv.Register("mul", func(ctx context.Context, params json.RawMessage) (interface{}, *jsonrpc2.Error) {
		nums, err := jsonrpc2.DecodeParams[[]float64](params)
		if err == nil {
			if len(nums) != 2 {
				return nil, jsonrpc2.ErrInvalidParams("you must input 2 params")
			}
			return roundToPrecision(nums[0]*nums[1], precision), nil
		}
		nums2, err := jsonrpc2.DecodeParams[ParmsXY](params)
		if nums2.X == nil || nums2.Y == nil {
			return nil, jsonrpc2.ErrInvalidParams("params object must contain x and y")
		}
		if err == nil {
			return roundToPrecision(*nums2.X**nums2.Y, precision), nil
		}
		return nil, jsonrpc2.ErrInvalidParams("incorrect params")
	})
	if err != nil {
		log.Fatal(err)
	}

	err = serv.Register("div", func(ctx context.Context, params json.RawMessage) (interface{}, *jsonrpc2.Error) {
		nums, err := jsonrpc2.DecodeParams[[]float64](params)
		if err == nil {
			if len(nums) != 2 {
				return nil, jsonrpc2.ErrInvalidParams("you must input 2 params")
			}
			if nums[1] == 0 {
				return nil, jsonrpc2.ErrInvalidParams("division by zero")
			}
			return roundToPrecision(nums[0]/nums[1], precision), nil
		}
		nums2, err := jsonrpc2.DecodeParams[ParmsXY](params)
		if nums2.X == nil || nums2.Y == nil {
			return nil, jsonrpc2.ErrInvalidParams("params object must contain x and y")
		}
		if err == nil {
			if *nums2.Y == 0 {
				return nil, jsonrpc2.ErrInvalidParams("division by zero")
			}
			return roundToPrecision((*nums2.X)/(*nums2.Y), precision), nil
		}
		return nil, jsonrpc2.ErrInvalidParams("incorrect params")
	})
	if err != nil {
		log.Fatal(err)
	}

	err = serv.Register("pre", func(ctx context.Context, params json.RawMessage) (interface{}, *jsonrpc2.Error) {
		prec, err := jsonrpc2.DecodeParams[ParmN](params)
		if err != nil {
			return nil, jsonrpc2.ErrInvalidParams("incorrect params")
		}
		if prec.N == nil {
			return nil, jsonrpc2.ErrInvalidParams("params object must contain N")
		}
		if *prec.N < 0 || *prec.N > 15 {
			return nil, jsonrpc2.ErrInvalidParams("N must be between 0 and 15")
		}
		precision = int(*prec.N)
		return "ok", nil
	})
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	regMethods()
	r := mux.NewRouter()

	r.Handle("/rpc", serv)

	log.Println("Server running on", 3000)
	if err := http.ListenAndServe("0.0.0.0:3000", r); err != nil {
		log.Fatal(err)
	}
}
