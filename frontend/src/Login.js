import React, { useState } from "react";
import './App.css'; // Import the custom CSS file
import Register from './Register.js'; // Register

const LoginPage = () => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');
  const [token, setToken] = useState('');

  const handleLogin = async (e) => {
    e.preventDefault();

    // In ra console để kiểm tra email và password
    console.log('Email:', email);
    console.log('Password:', password);

    try {
      const response = await fetch('http://localhost:8080/login', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ email, password }),
      });

      if (response.ok) {
        const data = await response.json();
        setToken(data.Token);  // Lưu lại token nếu đăng nhập thành công
        console.log('Login successful:', data); // In ra dữ liệu phản hồi từ API
      } else {
        console.error('Login failed:', response.statusText);
        setError('Đăng nhập thất bại. Vui lòng thử lại.');
      }
    } catch (err) {
      console.error('Error:', err);
      setError('Đăng nhập thất bại. Vui lòng thử lại.');
    }
  };

  return (
    <div className="container-fluid">
      <div className="row justify-content-center align-items-center vh-100">
        <div className="col-md-10 col-lg-8 d-md-flex box-login">
          {/* Left Side: Promotional Image */}
          <div className="col-md-6 d-flex justify-content-center align-items-center login-left">
            <div className="promotion-container">
              <img
                src="https://kdata.vn/kdata/images/banner/backgroud.png"
                alt=""
                className="img-ab"
              />
              <div className="slides-container">
                <div className="mySlides fadeslides active">
                  <img
                    src="https://kdata.vn/kdata/images/banner/banner_login_1.png"
                    className="slide-image"
                    alt=""
                  />
                </div>
                <div className="mySlides fadeslides">
                  <img
                    src="https://kdata.vn/kdata/images/banner/banner_login_2.png"
                    className="slide-image"
                    alt=""
                  />
                </div>
                <div className="mySlides fadeslides">
                  <img
                    src="https://kdata.vn/kdata/images/banner/banner_login_3.png"
                    className="slide-image"
                    alt=""
                  />
                </div>
                <div className="text-center mt-2">
                  <span className="dot active"></span>
                  <span className="dot"></span>
                  <span className="dot"></span>
                </div>
              </div>
            </div>
          </div>

          {/* Right Side: Login Form */}
          <div className="col-md-6 login-right d-flex justify-content-center align-items-center">
            <div className="form-container">
              <img
                src="https://kdata.vn/kdata/images/banner/Logo-Kdata-2000px 1.png"
                alt=""
                className="logo-login"
                width="140px"
              />
              <h1 className="login-title">Đăng nhập tài khoản</h1>
               <form onSubmit={handleLogin}> {/* Đổi sang onSubmit và dùng hàm handleLogin */}
      <div className="form-group mb-3">
        <label className="label-login" htmlFor="email">
          Email đăng nhập
        </label>
        <input
          id="email"
          type="email"
          className="form-control-login"
          placeholder="Email"
          name="email"
          required
          autoComplete="email"
          autoFocus
          value={email}
          onChange={(e) => setEmail(e.target.value)}
        />
      </div>
      <div className="form-group mb-3">
        <label className="label-login" htmlFor="password">
          Mật khẩu
        </label>
        <input
          id="password"
          type="password"
          className="form-control-login"
          placeholder="****************"
          name="password"
          required
          autoComplete="current-password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
        />
                </div>
                <div className="d-flex justify-content-between remember-forgot">
                  <div className="form-check">
                    <input
                      type="checkbox"
                      className="form-check-input"
                      name="remember"
                      id="remember"
                    />
                    <label className="form-check-label" htmlFor="remember">
                      Duy trì đăng nhập
                    </label>
                  </div>
                  <div>
                    <a href="https://kdata.vn/user/forget-password">Quên mật khẩu</a>
                  </div>
                </div>
                
                <div className="form-group">
                  <button
                    type="submit"
                    className="btn btn-primary submit px-4 py-2 w-100" onClick={handleLogin}
                  >
                    Đăng nhập
                  </button>
                </div>
                <div className="login-bottom mt-3 text-center">
                  <div>
                    <span>
                      Chưa có tài khoản?{' '}
                      <a href="Register" className="sign-up-now">
                        Đăng ký ngay
                      </a>
                    </span>
                  </div>
                  <div className="div-account">
                    <p className="account">
                      <small>Liên kết tài khoản</small>
                    </p>
                  </div>
                  <div className="login-with-social-network">
                    <div className="icon">
                      <a href="#">
                        <img
                          src="https://kdata.vn/kdata/images/banner/zalo.png"
                          alt=""
                          data-toggle="tooltip"
                          data-placement="top"
                          title="Zalo"
                        />
                      </a>
                      <a href="#">
                        <img
                          src="https://kdata.vn/kdata/images/banner/Telegram.png"
                          alt=""
                          data-toggle="tooltip"
                          data-placement="top"
                          title="Telegram"
                        />
                      </a>
                      <a
                        href="https://www.facebook.com/kdata.appota"
                        target="_blank"
                        rel="nofollow"
                      >
                        <img
                          src="https://kdata.vn/kdata/images/banner/Facebook.png"
                          alt=""
                          data-toggle="tooltip"
                          data-placement="top"
                          title="Facebook"
                        />
                      </a>
                    </div>
                    <div className="text-social mt-2">
                      <a href="">
                        Hướng dẫn liên kết tài khoản qua: Zalo, Telegram, Facebook
                      </a>
                    </div>
                  </div>
                </div>
              </form>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default LoginPage;
// import React, { useState } from 'react';

// const Login = () => {
//   const [email, setEmail] = useState('');
//   const [password, setPassword] = useState('');
//   const [error, setError] = useState('');
//   const [token, setToken] = useState('');

//   const handleLogin = async (e) => {
//     e.preventDefault();

//     // In ra console để kiểm tra email và password
//     console.log('Email:', email);
//     console.log('Password:', password);

//     try {
//       const response = await fetch('http://localhost:8080/login', {
//         method: 'POST',
//         headers: {
//           'Content-Type': 'application/json',
//         },
//         body: JSON.stringify({ email, password }),
//       });

//       if (response.ok) {
//         const data = await response.json();
//         setToken(data.Token);  // Lưu lại token nếu đăng nhập thành công
//         console.log('Login successful:', data); // In ra dữ liệu phản hồi từ API
//       } else {
//         console.error('Login failed:', response.statusText);
//         setError('Đăng nhập thất bại. Vui lòng thử lại.');
//       }
//     } catch (err) {
//       console.error('Error:', err);
//       setError('Đăng nhập thất bại. Vui lòng thử lại.');
//     }
//   };

//   return (
//     <form onSubmit={handleLogin}>
//       <div className="form-group mb-3">
//         <label className="label-login" htmlFor="email">Email đăng nhập</label>
//         <input
//           id="email"
//           type="email"
//           className="form-control-login"
//           placeholder="Email"
//           name="email"
//           required
//           autoComplete="email"
//           autoFocus
//           value={email}
//           onChange={(e) => setEmail(e.target.value)}
//         />
//       </div>
//       <div className="form-group mb-3">
//         <label className="label-login" htmlFor="password">Mật khẩu</label>
//         <input
//           id="password"
//           type="password"
//           className="form-control-login"
//           placeholder="****************"
//           name="password"
//           required
//           autoComplete="current-password"
//           value={password}
//           onChange={(e) => setPassword(e.target.value)}
//         />
//       </div>
//       <button type="submit" className="btn btn-primary" onClick={handleLogin}>  
//   Đăng nhập
// </button>

//       {error && <p style={{ color: 'red' }}>{error}</p>}
//     </form>
//   );
// };

// export default Login;
