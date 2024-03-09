window.onload = function() {
	const arr = ["This", "Is", "Rendered", "By", "RUST🦀"];
	let i = 0;
		let ele = document.querySelector("h1");
		setInterval(function () {
			ele.innerHTML = arr[i];
			i = (i + 1) % arr.length;
		}, 1000);
};
