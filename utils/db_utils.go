package utils

// import "context"

// func GetOne(ctx context.Context, query string, args ...interface{}) (map[string]interface{}, error) {
// 	rows, err := v.DB.QueryContext(ctx, query, args...)
// 	if err != nil {
// 		return nil, err
// 	}

// 	defer rows.Close()
// 	columns, _ := rows.Columns()
// 	values := make([]interface{}, len(columns))
// 	valuePtrs := make([]interface{}, len(columns))
// 	res := map[string]interface{}{}

// 	for rows.Next() {
// 		for i := range columns {
// 			valuePtrs[i] = &values[i]
// 		}

// 		rows.Scan(valuePtrs...)

// 		for i, col := range columns {
// 			val := values[i]

// 			b, ok := val.([]byte)
// 			var v interface{}
// 			if ok {
// 				v = string(b)
// 			} else {
// 				v = val
// 			}

// 			res[col] = v
// 		}
// 	}

// 	if err != nil {
// 		return map[string]interface{}{}, err
// 	}

// 	return res, nil
// }

// func Fetch(ctx context.Context, query string, args ...interface{}) ([]map[string]interface{}, error) {
// 	rows, err := v.DB.QueryContext(ctx, query, args...)
// 	if err != nil {
// 		return nil, err
// 	}

// 	defer rows.Close()
// 	columns, _ := rows.Columns()
// 	values := make([]interface{}, len(columns))
// 	valuePtrs := make([]interface{}, len(columns))
// 	result := []map[string]interface{}{}

// 	for rows.Next() {
// 		row := map[string]interface{}{}
// 		for i := range columns {
// 			valuePtrs[i] = &values[i]
// 		}

// 		rows.Scan(valuePtrs...)

// 		for i, col := range columns {
// 			val := values[i]

// 			b, ok := val.([]byte)
// 			var v interface{}
// 			if ok {
// 				v = string(b)
// 			} else {
// 				v = val
// 			}

// 			row[col] = v
// 		}
// 		result = append(result, row)
// 	}

// 	if err != nil {
// 		return []map[string]interface{}{}, err
// 	}

// 	return result, nil
// }
