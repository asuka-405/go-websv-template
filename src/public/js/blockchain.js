const blockChainContainer = document.querySelector(".blockchain-container")
const blocks = document.querySelectorAll(".blockchain-block")
const arrowCanvas = document.getElementById("arrowCanvas")

const dragdata = {
	item: null,
	offsetX: 0,
	offsetY: 0,
}

// Runtime calculations
const containerWidth = blockChainContainer.clientWidth
const firstBlock = document.querySelector(".blockchain-block")
const blockWidth = firstBlock ? firstBlock.offsetWidth : 80
const blockHeight = firstBlock ? firstBlock.offsetHeight : 100
const spacing = 50
const blockCountPerRow = Math.floor(
	(containerWidth - spacing) / (blockWidth + spacing)
)

let startX = spacing
let startY = spacing

const rowCount = Math.ceil(blocks.length / blockCountPerRow)
const newHeight =
	startY + (blockHeight + spacing) * (rowCount - 1) + blockHeight + spacing
blockChainContainer.style.height = `${newHeight}px`

// Function to draw arrows between blocks
function updateArrows() {
	arrowCanvas.innerHTML = "" // Clear existing arrows

	const containerRect = blockChainContainer.getBoundingClientRect()

	blocks.forEach((_, index) => {
		if (index === blocks.length - 1) return
		const block1 = blocks[index].getBoundingClientRect()
		const block2 = blocks[index + 1].getBoundingClientRect()

		const x1 = block1.right - containerRect.left
		const y1 = block1.top + block1.height / 2 - containerRect.top
		const x2 = block2.left - containerRect.left
		const y2 = block2.top + block2.height / 2 - containerRect.top

		const arrow = document.createElementNS("http://www.w3.org/2000/svg", "line")
		arrow.setAttribute("x1", x1)
		arrow.setAttribute("y1", y1)
		arrow.setAttribute("x2", x2)
		arrow.setAttribute("y2", y2)
		arrow.setAttribute("stroke", "black")
		arrow.setAttribute("stroke-width", "2")
		arrow.setAttribute("marker-end", "url(#arrowhead)")

		arrowCanvas.appendChild(arrow)
	})
}

// Add arrowhead marker
arrowCanvas.innerHTML = `
    <defs>
        <marker id="arrowhead" markerWidth="20" markerHeight="14" refX="10" refY="7" orient="auto">
            <polygon points="0 0, 20 7, 0 14" fill="white"/>
        </marker>
    </defs>
`

// Set initial positions of blocks with snake-like wrapping
blocks.forEach((block, index) => {
	const row = Math.floor(index / blockCountPerRow)
	const col = index % blockCountPerRow
	const isEvenRow = row % 2 === 0

	block.style.position = "absolute"
	block.style.cursor = "grab"
	block.style.left = isEvenRow
		? `${startX + (blockWidth + spacing) * col}px`
		: `${startX + (blockWidth + spacing) * (blockCountPerRow - 1 - col)}px`
	block.style.top = `${startY + (blockHeight + spacing) * row}px`
})

blocks.forEach((block) => {
	block.addEventListener("mousedown", (e) => {
		const blockRect = block.getBoundingClientRect()
		dragdata.item = block
		dragdata.offsetX = e.clientX - blockRect.left
		dragdata.offsetY = e.clientY - blockRect.top
		dragdata.item.style.cursor = "grabbing"
		dragdata.item.style.zIndex = "1000"
		e.target.style.userSelect = "none"
	})
})

document.addEventListener("mousemove", (e) => {
	if (!dragdata.item) return

	let x = e.clientX - dragdata.offsetX
	let y = e.clientY - dragdata.offsetY

	const containerRect = blockChainContainer.getBoundingClientRect()
	const blockRect = dragdata.item.getBoundingClientRect()

	const minX = containerRect.left
	const minY = containerRect.top
	const maxX = containerRect.right - blockRect.width
	const maxY = containerRect.bottom - blockRect.height

	// Ensure block stays within container
	x = Math.max(minX, Math.min(x, maxX))
	y = Math.max(minY, Math.min(y, maxY))

	// Convert to relative position inside container
	dragdata.item.style.left = `${x - containerRect.left}px`
	dragdata.item.style.top = `${y - containerRect.top}px`

	updateArrows() // Update arrows on drag
})

document.addEventListener("mouseup", () => {
	if (dragdata.item) dragdata.item.style.cursor = "grab"
	dragdata.item.style.zIndex = ""
	dragdata.item = null
})

// Initial arrow rendering
updateArrows()
