package quantum

func NewMatrix(rows, cols int) Matrix {
	data := make([][]complex128, rows)
	for i := 0; i < rows; i++ {
		data[i] = make([]complex128, cols)
	}
	return Matrix{
		Rows: rows,
		Cols: cols,
		Data: data,
	}
}

func (m *Matrix) CanMultiply(n *Matrix) bool {
	return m.Cols == n.Rows
}

func (m *Matrix) MustMultiply(n *Matrix) Matrix {
	if !m.CanMultiply(n) {
		panic("cannot multiply matrices")
	}
	result := NewMatrix(m.Rows, n.Cols)
	for i := 0; i < m.Rows; i++ {
		for j := 0; j < n.Cols; j++ {
			var sum complex128
			for k := 0; k < m.Cols; k++ {
				sum += m.Data[i][k] * n.Data[k][j]
			}
			result.Data[i][j] = sum
		}
	}
	return result
}

func tensorGateMatrix(a, b *Matrix) Matrix {
	result := NewMatrix(a.Rows*b.Rows, a.Cols*b.Cols)
	for i := 0; i < a.Rows; i++ {
		for j := 0; j < a.Cols; j++ {
			for k := 0; k < b.Rows; k++ {
				for l := 0; l < b.Cols; l++ {
					result.Data[i*b.Rows+k][j*b.Cols+l] = a.Data[i][j] * b.Data[k][l]
				}
			}
		}
	}
	return result
}
