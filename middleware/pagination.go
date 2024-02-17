package middleware

// func PaginationMiddleware(c *fiber.Ctx) error {
// 	var err error

// 	offset := 0∫
// 	limit := 10

// 	q := c.Queries()

// 	if len(q["page[offset]"]) > 0 {
// 		offset, err = strconv.Atoi(q["page[offset]"])
// 		if err != nil {
// 			return c.Status(400).JSON("Bad pagination")
// 		∫}
// 	}

// 	if len(q["page[limit]"]) > 0 {
// 		limit, err = strconv.Atoi(q["page[limit]"])
// 		if err != nil {
// 			return c.Status(400).JSON("Bad pagination")
// 		}
// 	}

// 	c.Locals("offset", offset)
// 	c.Locals("limit", limit)

// 	return c.Next()
// }
