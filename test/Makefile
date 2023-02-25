.PHONY: run
run:
	docker run --rm \
		--name telegraf-pihole \
		-v $(PWD)/test/pihole-FTL.db:/etc/pihole/pihole-FTL.db:ro \
		-v $(PWD)/test/telegraf.conf:/etc/telegraf/telegraf.conf \
		-v $(PWD)/test/pihole.conf:/etc/telegraf/pihole.conf \
		ghcr.io/wyvernzora/telegraf-pihole:dev
