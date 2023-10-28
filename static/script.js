  function openDepositPopup() {
    var depositPopup = document.getElementById("deposit-popup");
    depositPopup.style.display = "block";

  }

  function closeDepositPopup() {
    var depositPopup = document.getElementById("deposit-popup");
    depositPopup.style.display = "none";
  }
  
  document.addEventListener('DOMContentLoaded', () => {
    const spinButtons = document.querySelectorAll('.spin-button');
    const rouletteWheel = document.getElementById('roulette-wheel');
    const countdownDisplay = document.getElementById('countdown-display');
    const betForm = document.getElementById('bet-form');
  
    let countdown; 
  
    spinButtons.forEach(spinButton => {
      spinButton.addEventListener('click', () => {
  
        clearInterval(countdown);
  
        rouletteWheel.classList.add('spinning');
        
  
     
        fetch('/spin')
          .then(response => response.text())
          .then(winningDegrees => {
            console.log('Winning Degrees:', winningDegrees);
            console.log('Hello!');
  
         
            rouletteWheel.style.transition = 'transform 5s ease-in-out';
            rouletteWheel.style.transform = `rotate(${winningDegrees}deg)`;
  
            
            setTimeout(() => {
              rouletteWheel.style.transition = 'transform 0.00001s ease-in-out'; 
              rouletteWheel.style.transform = 'rotate(0deg)';
            }, 7000); 
  
            
            startCountdown(7); 
          })
          .catch(error => {
            console.error('Error:', error);
          });
      });
    });
  t
    
    function startCountdown(seconds) {
      let remainingSeconds = seconds;
      countdownDisplay.textContent = 'Betting for next round open for ' + remainingSeconds + ' seconds';
  
      countdown = setInterval(() => {
        remainingSeconds--;
        countdownDisplay.textContent = 'Betting for next round open for ' + remainingSeconds + ' seconds';
  
        if (remainingSeconds <= 0) {
          clearInterval(countdown); 
          countdownDisplay.textContent = 'Betting closed';
        }
      }, 1000);
    }
  
    
    setInterval(() => {
      
      spinButtons[0].click(); 
    }, 10000);
  });