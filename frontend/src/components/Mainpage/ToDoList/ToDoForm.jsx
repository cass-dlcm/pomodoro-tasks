import React, { useState, useEffect, useRef } from 'react';

function ToDoForm(props){
  const [input, setInput] = useState(props.edit ? props.edit.value : '');

  const inputRef = useRef(null);

  useEffect(() => {
    inputRef.current.focus();
  });

  const handleChange = e => {
    setInput(e.target.value);
  };

  const handleSubmit = e => {
    e.preventDefault();

    props.onSubmit({
      id: Math.floor(Math.random() * 10000),
      text: input
    });
    setInput('');
  };

  return (
    <form onSubmit={handleSubmit} className='todo-form'>
          <input
            placeholder='Add a To Do'
            value={input}
            onChange={handleChange}
            name='text'
            className='todo-input'
            ref={inputRef}
          />
          <button 
            onClick={handleSubmit} 
            className='todo-button'>
            Add Task
          </button>
        
    </form>
  );
}
export default ToDoForm;
