const _service = {
    refresh: function() {
        console.log('Refreshing')
        setLoading(true)
        const container = document.getElementById('container')
        const table = container.querySelector('table')
        console.log('table', table)
        table.remove()
        const data = _api.getAll()
        console.log('Received', data)
        buildTable(container, data)
        setLoading(false)
    }
}

function setLoading (loading) {
    const el = document.getElementById('loading-indicator')
    if (loading)
        el.classList.add('active')
    else
        el.classList.remove('active')
}

function buildTable(parent, data) {
    const table = document.createElement('table')
    const headRow = document.createElement('tr')
    const headerNumber = document.createElement('th')
    headerNumber.appendChild(document.createTextNode('Number'))
    headRow.appendChild(headerNumber)
    const headerTask = document.createElement('th')
    headerTask.appendChild(document.createTextNode('Task'))
    headRow.appendChild(headerTask)
    headRow.appendChild(document.createElement('th'))
    table.appendChild(headRow)
    for (var index = 0; index < data.Length; index++) {
        const bodyRow = document.createElement('tr')
        const cellNumber = document.createElement('td')
        const cellNumberSpan = document.createElement('span')
        cellNumberSpan.appendChild(document.createTextNode((index + 1)))
        const checkbox = document.createElement('input')
        checkbox.setAttribute('type', 'checkbox')
        checkbox.setAttribute('value', index)
        cellNumber.appendChild(cellNumberSpan)
        cellNumberSpan.appendChild(checkbox)
        bodyRow.appendChild(cellNumber)
        table.appendChild(bodyRow)
    }
    parent.appendChild(table)

}