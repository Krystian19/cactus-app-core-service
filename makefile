run :
	watchexec --exts rs -i target --restart "cargo run"

gen : SHELL:=/bin/bash
gen :
	./proto-gen.sh

build :
	cargo build --release