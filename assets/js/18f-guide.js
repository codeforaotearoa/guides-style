$(function() {
  $('.sidebar-nav').each(function() {
    accordion($(this));
  });
});

// http://bryanbraun.github.io/anchorjs/
anchors.options = {
  visible: 'touch'
};
anchors.add('h3,h4,h5');
