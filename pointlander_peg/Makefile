all: install_peg gen_hour_min

install_peg:
	which peg >/dev/null || go get -u github.com/pointlander/peg

gen_hour_min: hour_min/hour_min.peg
	peg hour_min/hour_min.peg
