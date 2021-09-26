import { CountdownCircleTimer } from 'react-countdown-circle-timer'
/*
const children = ({ remainingTime }) => {
  const hours = Math.floor(remainingTime / 3600)
  const minutes = Math.floor((remainingTime % 3600) / 60)
  const seconds = remainingTime % 60

  return `${hours}:${minutes}:${seconds}`
}*/
export const Clock = (size) => {
  return(
    <div class="countdown">
      <CountdownCircleTimer
        size="300"
        isPlaying
        duration={100}

        colors={[
          ['#004777', 0.33],
          ['#F7B801', 0.33],
          ['#A30000', 0.33],]}
        >
       {({ remainingTime }) => remainingTime}
     
      </CountdownCircleTimer>

  
    </div>
  );
}