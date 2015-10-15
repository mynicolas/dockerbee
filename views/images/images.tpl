<docker-main title="Images">
	<docker-menu class="menu"></docker-menu>
	$$range .Imgs**
		<docker-images 
			class="detail" 
			image_id="$$.ID**" 
			repo_tags="$$.RepoTags**" 
			created_time="$$.Created**"  
			size="$$.Size**" 
			virtual_size="$$.VirtualSize**" 
			parent_id="$$.ParentID**" 
			repo_digests="$$.RepoDigests**" 
			labels="$$.Labels**" 
		></docker-images>
	$$end**
</docker-main>