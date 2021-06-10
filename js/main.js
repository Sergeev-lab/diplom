

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

  var app = new Vue ({
    delimiters: ['${', '}$'],
    el: '#auth',
    data: {
      auth: '',
    },
    methods: {
      getToken: function() {
        var req = new XMLHttpRequest();
        req.open('GET', 'http://localhost:8181/', false);
        req.send(null);
        var headers = req.getResponseHeader('Authorization');
        this.auth = headers;
        return;
      },
    },
    created: function() {
      this.getToken();
    },
  });

  var ap = new Vue ({
    delimiters: ['${', '}$'],
    el: '#test',
    data: {
      dan: app.auth,
      text: false,
    },
  });
  var x = 0;
  var a = new Vue ({
    delimiters: ['${', '}$'],
    el: '#myData',
    data: {
      file: '',
      click: false,
      name: '',
      btnText: 'Изменить',
      visible: true,
    },
    methods: {
      previewFiles: function() {
        this.file = 'has-file';
        return
       },
      clickFunc: function() {
        this.click = !this.click;
        if (this.click == true) {
          this.btnText = 'Сохранить';
        } else {
          this.btnText = 'Изменить';
        };
      },
      clickBtn: function() {
        window.location.reload();
      },
    },
    created: function() {
      if (document.location.pathname == '/user/') {
        var oldname = document.getElementsByName("newName")[0];
        this.name = oldname.value;
      }
    },
  });