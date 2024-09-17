import React, { useState } from 'react';
import './App.css'; // Import CSS riêng


const RegisterPage = () => {
   // State để lưu trữ các giá trị của form
   const [email, setEmail] = useState('');
   const [password, setPassword] = useState('');
   const [passwordConfirmation, setPasswordConfirmation] = useState('');
   const [name, setName] = useState('');
   const [phone, setPhone] = useState('');
   const [error, setError] = useState('');
   const [success, setSuccess] = useState('');
 
   // Hàm xử lý khi form được gửi
   const handleSubmit = async (e) => {
     e.preventDefault();
     console.log('Email:', email);
     console.log('Password:', password);
     console.log('name:', name);
     console.log('phone:', phone);
     // Kiểm tra tính hợp lệ của mật khẩu
     if (password !== passwordConfirmation) {
       setError('Mật khẩu và xác nhận mật khẩu không khớp');
       return;
     }
 
     // Gửi dữ liệu đến API
     try {
       const response = await fetch('http://localhost:8080/register', {
         method: 'POST',
         headers: {
           'Content-Type': 'application/json',
         },
         body: JSON.stringify({
           email: email,
           username: name, // Thay đổi từ 'name' thành 'username'
           phone: phone,
           pass: password,
         }),
       });
 
       if (!response.ok) {
         const data = await response.json();
         setError(data.Message || 'Đăng ký thất bại');
         return;
       }
 
       const data = await response.json();
       console.log('đăng kí thành công:', data);

       setSuccess(data.Message || 'Đăng ký thành công');
       setError('');
     } catch (err) {
       setError('Lỗi kết nối đến máy chủ');
     }
   };

  return (
    
    <div
      style={{
        background: 'url(https://kdata.vn/kdata/images/bg-login.jpg)',
        backgroundSize: 'cover',
        backgroundRepeat: 'no-repeat',
        minHeight: '100vh',
        display: 'flex',
        alignItems: 'center',
      }}
    >
      <div id="loader_spinner"  />
      <section className="ftco-section login">
        <div className="container">
          <div className="row justify-content-center">
            <div className="col-md-7 col-lg-6">
              <div className="d-md-flex box-login">
                <div className="col-md-12">
                  <div className="login-right">
                    <div className="logo-main">
                      <img
                        src="https://kdata.vn/kdata/images/banner/Logo-Kdata-2000px 1.png"
                        alt="KData Logo"
                        className="logo-login"
                      />
                    </div>
                    <h1 className="login-title">Đăng ký tài khoản</h1>
                    {error && <div className="error">{error}</div>}
      {success && <div className="success">{success}</div>}
      <form onSubmit={handleSubmit}>
        <div className="form-group mb-3">
          <label className="label-login" htmlFor="email">
            Email
          </label>
          <input
            id="email"
            type="email"
            className="form-control-login"
            placeholder="Email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            required
            autoComplete="email"
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
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            required
            autoComplete="new-password"
          />
        </div>
        <div className="form-group mb-3">
          <label className="label-login" htmlFor="password-confirm">
            Nhập lại mật khẩu
          </label>
          <input
            id="password-confirm"
            type="password"
            className="form-control-login"
            placeholder="****************"
            value={passwordConfirmation}
            onChange={(e) => setPasswordConfirmation(e.target.value)}
            required
            autoComplete="new-password"
          />
        </div>
        <div className="row">
          <div className="col-md-6">
            <div className="form-group mb-3">
              <label className="label-login" htmlFor="name">
                Tên đầy đủ
              </label>
              <input
                id="name"
                type="text"
                className="form-control-login"
                placeholder="Nhập tên"
                value={name}
                onChange={(e) => setName(e.target.value)}
                required
                autoFocus
              />
            </div>
          </div>
          <div className="col-md-6">
            <div className="form-group mb-3">
              <label className="label-login" htmlFor="phone">
                Điện thoại
              </label>
              <input
                id="phone"
                type="text" // Changed from 'number' to 'text' to handle formatting
                className="form-control-login"
                placeholder="Nhập số điện thoại"
                value={phone}
                onChange={(e) => setPhone(e.target.value)}
                required
                maxLength="11"
              />
                          </div>
                        </div>
                      </div>
                      <div className="form-group mb-3">
                        <div className="d-md-flex remember-forgot">
                          <div className="w-50 text-left" style={{ overflow: 'hidden' }}>
                            <div className="input-group">
                              <div className="g-recaptcha" data-sitekey="6LeWx2ogAAAAABZ5B3ZlZbdeJOzIGhmPyo2zsW8d" />
                            </div>
                          </div>
                        </div>
                      </div>
                      <div style={{ margin: '15px 0', fontSize: '14px', textAlign: 'center' }}>
                        <p>
                          Khi click nút đăng ký, nghĩa là bạn đã đồng ý với{' '}
                          <a href="https://kdata.vn/general-rules" target="_blank" rel="noopener noreferrer">
                            quy định chung
                          </a>{' '}
                          và{' '}
                          <a
                            href="https://kdata.vn/information-security-policy"
                            target="_blank"
                            rel="noopener noreferrer"
                          >
                            chính sách bảo mật
                          </a>{' '}
                          của chúng tôi.
                        </p>
                      </div>
                      <div className="form-group">
                        <button type="submit" className="form-control-btn btn btn-primary submit px-3">
                          Đăng ký
                        </button>
                      </div>
                      <div className="login-bottom">
                        <div>
                          <span>
                            Đã có tài khoản?{' '}
                            <a href="https://kdata.vn/user/login" className="sign-up-now">
                              Đăng nhập
                            </a>
                          </span>
                        </div>
                      </div>
                    </form>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </section>
    </div>
  );
};

export default RegisterPage;
// import React, { useState } from 'react';

// const Register = () => {
//   // State để lưu trữ các giá trị của form
//   const [email, setEmail] = useState('');
//   const [password, setPassword] = useState('');
//   const [passwordConfirmation, setPasswordConfirmation] = useState('');
//   const [name, setName] = useState('');
//   const [phone, setPhone] = useState('');
//   const [error, setError] = useState('');
//   const [success, setSuccess] = useState('');

//   // Hàm xử lý khi form được gửi
//   const handleSubmit = async (e) => {
//     e.preventDefault();
//     console.log('Email:', email);
//     console.log('Password:', password);
//     console.log('name:', name);
//     console.log('phone:', phone);
//     // Kiểm tra tính hợp lệ của mật khẩu
//     if (password !== passwordConfirmation) {
//       setError('Mật khẩu và xác nhận mật khẩu không khớp');
//       return;
//     }

//     // Gửi dữ liệu đến API
//     try {
//       const response = await fetch('http://localhost:8080/register', {
//         method: 'POST',
//         headers: {
//           'Content-Type': 'application/json',
//         },
//         body: JSON.stringify({
//           email: email,
//           username: name, // Thay đổi từ 'name' thành 'username'
//           phone: phone,
//           pass: password,
//         }),
//       });

//       if (!response.ok) {
//         const data = await response.json();
//         setError(data.Message || 'Đăng ký thất bại');
//         return;
//       }

//       const data = await response.json();
//       setSuccess(data.Message || 'Đăng ký thành công');
//       setError('');
//     } catch (err) {
//       setError('Lỗi kết nối đến máy chủ');
//     }
//   };

//   return (
//     <div>
//       <h2>Đăng ký</h2>
//       {error && <div className="error">{error}</div>}
//       {success && <div className="success">{success}</div>}
//       <form onSubmit={handleSubmit}>
//         <div className="form-group mb-3">
//           <label className="label-login" htmlFor="email">
//             Email
//           </label>
//           <input
//             id="email"
//             type="email"
//             className="form-control-login"
//             placeholder="Email"
//             value={email}
//             onChange={(e) => setEmail(e.target.value)}
//             required
//             autoComplete="email"
//           />
//         </div>
//         <div className="form-group mb-3">
//           <label className="label-login" htmlFor="password">
//             Mật khẩu
//           </label>
//           <input
//             id="password"
//             type="password"
//             className="form-control-login"
//             placeholder="****************"
//             value={password}
//             onChange={(e) => setPassword(e.target.value)}
//             required
//             autoComplete="new-password"
//           />
//         </div>
//         <div className="form-group mb-3">
//           <label className="label-login" htmlFor="password-confirm">
//             Nhập lại mật khẩu
//           </label>
//           <input
//             id="password-confirm"
//             type="password"
//             className="form-control-login"
//             placeholder="****************"
//             value={passwordConfirmation}
//             onChange={(e) => setPasswordConfirmation(e.target.value)}
//             required
//             autoComplete="new-password"
//           />
//         </div>
//         <div className="row">
//           <div className="col-md-6">
//             <div className="form-group mb-3">
//               <label className="label-login" htmlFor="name">
//                 Tên đầy đủ
//               </label>
//               <input
//                 id="name"
//                 type="text"
//                 className="form-control-login"
//                 placeholder="Nhập tên"
//                 value={name}
//                 onChange={(e) => setName(e.target.value)}
//                 required
//                 autoFocus
//               />
//             </div>
//           </div>
//           <div className="col-md-6">
//             <div className="form-group mb-3">
//               <label className="label-login" htmlFor="phone">
//                 Điện thoại
//               </label>
//               <input
//                 id="phone"
//                 type="text" // Changed from 'number' to 'text' to handle formatting
//                 className="form-control-login"
//                 placeholder="Nhập số điện thoại"
//                 value={phone}
//                 onChange={(e) => setPhone(e.target.value)}
//                 required
//                 maxLength="11"
//               />
//             </div>
//           </div>
//         </div>
//         <button type="submit" className="btn btn-primary">Đăng ký</button>
//       </form>
//     </div>
//   );
// };

// export default Register;
