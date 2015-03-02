$(document).ready(function(e) {
	index = 0;
	$.each(color_data, function(key, value) {
		console.log(key, value);
		row = $('<div>').text(value.Name).addClass('language').css('backgroundColor', value.Color.Hex);
		$('#color-list').append(row);
	});
});