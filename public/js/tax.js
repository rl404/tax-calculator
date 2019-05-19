$(document).ready(function() {
	$('#add-more-tax').click(function(e) {
		e.preventDefault()
		rowCount = $('#tax-table').find('.tax-row').size()
		newRow = $('#empty-tax-form').clone().insertBefore($('#add-more-row')).show()
		newRow.find('td:first').html(rowCount+'.')
	})
})