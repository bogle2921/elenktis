import React, { useState, useEffect, useRef } from 'react';
import './Login.css'

//const USERNAME_REGEX = /^[a-zA-Z][a-zA-Z0-9-_]{3,23}$/;
//const PASS_REGEX = /^(?=.*[a-z])(?=.*[A-Z])(?=.*[0-9])(?=.*[!@#$%]).{8,24}$/;

const Login = () => {

    const userRef = useRef();
    const errRef = useRef();

    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');

    const [errMsg, setErrMsg] = useState('');
    const [success, setSuccess] = useState(false);

    useEffect(() => {
        userRef.current.focus();
    }, []);

    useEffect(() =>{
        setErrMsg('');
    }, [username, password]);

    const handleSubmit = async (e) => {
        e.preventDefault();
        fetch('http://localhost:8080', {
            method: 'POST',
            body: {
                username: username,
                password: password
            }
        }).then(res => {
            return res.json();
        }).catch(err => console.log(`Error: ${err}`));
    }

    return (
        <div className='container'>
            <div classname='header'>
                <div className='text'>Login</div>
                <div className='underline'></div>
            </div>
            <form className='inputs' onSubmit={handleSubmit}>
                <label htmlFor='username'>Username:</label>
                <div className='input'>
                    <input type='text' id='username' ref={userRef} onChange={(e) => setUsername(e.target.value)} value={username} required/>
                </div>
                <label htmlFor='password'>Password:</label>
                <div className='input'>
                    <input type='password' id='password' onChange={(e) => setPassword(e.target.value)} value={password} required/>
                </div>
                <div className='.submit-container'>
                    <button className='submit'>Sign In</button>
                </div>
            </form>
        </div>
    )
}

export default Login