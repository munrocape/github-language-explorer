$(document).ready(function(e) {
	index = 0;
	$.each(color_data, function(key, value) {
		console.log(key, value);
		row = $('<div>').text(value.Name);
		row.css('backgroundColor', value.Color.Hex);
		row.addClass('language');
		$('#color-list').append(row);
	});
});
