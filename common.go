package fourier

func conjugate(c complex128) complex128 {
	return complex(real(c), -imag(c))
}
