// MIT License
//
// Copyright (c) 2021 SundaeSwap Finance
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package gql

import (
	"context"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

const delay = time.Millisecond * 1500

type TxFeeArgs struct {
	Raw       string
	TxIn      int32
	TxOut     int32
	Witnesses int32
}

func (r *Resolver) TxFee(ctx context.Context, args TxFeeArgs) (string, error) {
	data, err := base64.StdEncoding.DecodeString(args.Raw)
	if err != nil {
		return "", fmt.Errorf("failed to calculate fee: %w", err)
	}

	f, err := ioutil.TempFile(filepath.Join(r.config.CLI.Dir, "/tmp"), "script")
	if err != nil {
		return "", fmt.Errorf("failed to calculate fee: %w", err)
	}
	defer os.Remove(f.Name())
	defer f.Close()

	if _, err := f.Write(data); err != nil {
		return "", fmt.Errorf("failed to calculate fee: %w", err)
	}

	return r.config.CLI.MinFee(ctx, f.Name(), args.TxIn, args.TxOut, args.Witnesses)
}
