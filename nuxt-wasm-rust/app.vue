<script setup>

import init, { greets, add } from "./pkg/wasmfile.js"

const greeting = ref("");
const input1 = ref(0);
const input2 = ref(0);
const result = ref(0);

onBeforeMount(() => {
	init().then(() => {
		console.log("WASM initialized!");
		greeting.value = greets();
	});	
	
});

function using_rust_add(a, b) {
	result.value = add(a, b)
}

</script>

<template>
	<main>
		<h1>{{ greeting }}</h1>
		<div class="loading"></div>

		<div class="set">
			<h1 v-if="!result">add 2 numbers</h1>
			<h1 v-else>{{ result }}</h1>
			<input type="number" v-model="input1"/>
			<input type="number" v-model="input2"/>
			<button @click="using_rust_add(input1, input2)"> ADD </button>
		</div>
		
	</main>
</template>