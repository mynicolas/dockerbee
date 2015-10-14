<docker-main title="Containers">
	<docker-menu class="menu"></docker-menu>
	$$range .Conts**
		<docker-containers class="detail" container_name="$$.Name**" container_status="$$.Status**"></docker-containers>
	$$end**
</docker-main>