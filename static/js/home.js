document.addEventListener('DOMContentLoaded', function() {
  const slides = document.querySelectorAll('.hero-slide');
  const prevBtn = document.querySelector('.slide-prev');
  const nextBtn = document.querySelector('.slide-next');
  let currentIndex = 0;

  function showSlide(index) {
    slides.forEach((slide, i) => {
      if (i == index) {
        slide.style.display = 'block';
      } else {
        slide.style.display = 'none';
      }
    });
  }

  function prevSlide() {
    currentIndex = (currentIndex - 1 + slides.length) % slides.length;
    showSlide(currentIndex);
  }

  function nextSlide() {
    currentIndex = (currentIndex + 1) % slides.length;
    showSlide(currentIndex);
  }

  prevBtn.addEventListener('click', prevSlide);
  nextBtn.addEventListener('click', nextSlide);

  showSlide(currentIndex);
});