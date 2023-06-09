// Copyright 2023 The FHub Authors
// This file is part of FHub
//
// This file is part of FHub.
// FHub is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// FHub is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with FHub. If not, see <https://www.gnu.org/licenses/>.

name: string
specVersion: "0.1"
version: string
build: {
  local?: {
    source: string | *"./"
  }
  container?: {
    image?: string
    containerFile?: string
    source: string | *"/app"
  }
}
serving: {
  http?: {
    url: string
  }
  grpc?: {}
}
functions: {
  [string]: {
    inputs: {
      [string]: number | string | bool | null
    }
    outputs: {
      [string]: number | string | bool | null
    }
  }
}