/*
   main.go
   Copyright (C) 2022 gltile

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU Affero General Public License as
   published by the Free Software Foundation, either version 3 of the
   License, or (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU Affero General Public License for more details.

   You should have received a copy of the GNU Affero General Public License
   along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package main

import (
	"asciichat/proto"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		c.Accepts("application/vnd.google.protobuf")
		c.Accepts("application/x-protobuf")
		c.Accepts("application/protobuf")

		body := &proto.RepeatRequest{}

		if err := proto.Unmarshal(c); err != nil {
			log.Println("Error while unmarshalling protobuf")
		}

	})

	log.Fatal(app.Listen("0.0.0.0:8080"))
}
