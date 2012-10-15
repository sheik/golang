/*global $:false, prettyPrint:false */

$(function () {
  var opts =
      { container: 'body'
      , file: { defaultContent: "#GoWiki\nThis is some default content. Go ahead, _change me_!" }
      , focusOnLoad: true
	  , clientSideStorage: false
	  , basePath: 'static'
      }
    , editor = new EpicEditor(opts).load();
    
  // So people can play with it in their console
  window.editor = editor;
  window.example = example;
  
  $('pre').addClass('prettyprint')
  prettyPrint()

  $(window).resize(function () {
    $('#toc').height(window.innerHeight + 'px');
  }).trigger('resize');
});
