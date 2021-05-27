var swiper = new Swiper(".mySwiper", {
    speed: 600,
    parallax: true,
    spaceBetween: 30,
    centeredSlides: true,
    autoplay: {
        delay: 3500,
        disableOnInteraction: false,
    },
    pagination: {
      el: ".swiper-pagination",
      clickable: true,
    },
    navigation: {
      nextEl: ".swiper-button-next",
      prevEl: ".swiper-button-prev",
    },
  });

  var app = new Vue({
    el: '#all',
    data: {
      visible: false,
    },
  });