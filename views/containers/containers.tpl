<docker-main title="Containers">
	<docker-menu class="menu"></docker-menu>
	$$range .Conts**
		<docker-containers 
			class="detail" 
			container_id="$$.ID**" 
			image="$$.Image**" 
			command="$$.Command**" 
			created_time="$$.Created**" 
			ports="$$.Ports**" 
			size="$$.SizeRw**" 
			root_size="$$.SizeRootFs**" 
			label="$$.Labels**" 
			name="$$.Names**" 
			status="$$.Status**" 
		></docker-containers>
	$$end**
</docker-main>