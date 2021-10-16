import { CountdownCircleTimer } from 'react-countdown-circle-timer'
import { useState } from 'react';
import React from 'react';

export const BreakClock  = (size, isPlaying) => {
    
    const [key, setKey] = useState(0);

    const children = ({ remainingTime }) => {
    const minutes = Math.floor(remainingTime / 60)
    const seconds = remainingTime % 60
        return <div style={{fontSize: "xxx-large"}}>{`${minutes}:${seconds}`}</div>
  }

  return(
    <div class="clock">
      <CountdownCircleTimer
        key={key}
        isPlaying 
        size="300"
        duration={300}

        colors={[
          ['#004777', 0.33],
          ['#F7B801', 0.33],
          ['#A30000', 0.33],]}
        >

       {children}

      </CountdownCircleTimer>
  
    </div>
  );
  
}
