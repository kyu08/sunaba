.PHONY: init-ubuntu
init-ubuntu:
	su \
	&& sudo chmod 666 /var/run/docker.sock
	&& bash ./setup_mac.sh
	&& bash ./symlink.sh
	&& bash ./check_mac.sh
	&& tmux

.PHONY: start
start:
	multipass start UBUNTU

.PHONY: into
into:
	multipass shell UBUNTU

.PHONY: stop
stop:
	multipass stop UBUNTU

.PHONY: dotfiles
dotfiles:
	./symlink.sh

.PHONY: tinet-upconf
tinet-upconf:
	tinet upconf -c /mnt/c/tinet/spec_${arg}.yaml | sh -x

.PHONY: tinet-up
tinet-up:
	tinet up -c /mnt/c/tinet/spec_${arg}.yaml | sh -x

.PHONY: tinet-conf
tinet-conf:
	tinet conf -c /mnt/c/tinet/spec_${arg}.yaml | sh -x

.PHONY: tinet-test
tinet-test:
	tinet test -c /mnt/c/tinet/spec_${arg}.yaml | sh -x

.PHONY: tinet-down
tinet-down:
	tinet down -c /mnt/c/tinet/spec_${arg}.yaml | sh -x
