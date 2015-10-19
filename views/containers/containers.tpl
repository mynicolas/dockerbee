<dom-module id="container-content">
	<template>
		<style>

		</style>
		<docker-dialog-create id="create_dialog"></docker-dialog-create>
		<docker-main title="Containers">
			<paper-button raised class="tool" on-click="create">Create</paper-button>
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
	</template>
	<script>
	Polymer({
		is: "container-content",
		create: function(e){
			this.$.create_dialog.open();
		}
	});
	</script>
</dom-module>

<container-content></container-content>
