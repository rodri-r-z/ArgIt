/*
   The ArgIt Project
   -----------------------------------------------------
   This code is released under the GNU GPL v3 license.
   For more information, please visit:
   https://www.gnu.org/licenses/gpl-3.0.html
   -----------------------------------------------------
   Copyright (c) 2025 Rodrigo R. & All Contributors
   This program comes with ABSOLUTELY NO WARRANTY.
   For details type `show w`. This is free software,
   and you are welcome to redistribute it under certain
   conditions; type `show c` for details.
*/

package value

func GetOrFallback[T comparable](value *interface{}, fallback T) T {
	if value == nil {
		return fallback
	}

	// Try to convert the value to T
	val, castOk := (*value).(T)
	if !castOk {
		return fallback
	}

	return val
}

func ConvertBool(value *interface{}) bool {
	return GetOrFallback(value, false)
}

func ConvertString(value *interface{}) string {
	return GetOrFallback(value, "")
}

func ConvertInt(value *interface{}) int {
	return GetOrFallback(value, 0)
}

func ConvertFloat(value *interface{}) float64 {
	return GetOrFallback(value, 0.0)
}
