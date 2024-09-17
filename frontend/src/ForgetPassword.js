import React from 'react';
import './App.css';

const  ForgetPassword = () => {
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
      <div id="loader_spinner" style={{ display: 'none' }}></div>

      <section className="ftco-section login">
        <div className="container">
          <div className="row justify-content-center">
            <div className="col-md-7 col-lg-5">
              <div className="d-md-flex box-login">
                <div className="col-md-12">
                  <div className="login-right">
                    <div className="logo-main">
                      <img
                        src="https://kdata.vn/kdata/images/banner/Logo-Kdata-2000px 1.png"
                        alt="Logo"
                        className="logo-login"
                      />
                    </div>
                    <h1 className="login-title">Lấy lại mật khẩu</h1>

                    <form
                      method="POST"
                      action="https://kdata.vn/user/forget-password"
                    >
                      <input
                        type="hidden"
                        name="_token"
                        value="JPnA2SpeUx1TrbltG8TYGoGTj6FvvxLtmLI8YGZJ"
                      />
                      <div className="col-lg-12 mes_info"></div>
                      <div className="form-group mb-3">
                        <label className="label-login" htmlFor="email">
                          Email khôi phục
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
                        />
                      </div>
                      <div className="form-group mb-3">
                        <div className="d-md-flex remember-forgot">
                          <div className="w-50 text-left" style={{ overflow: 'hidden' }}>
                            <div className="input-group">
                              <div
                                className="g-recaptcha"
                                data-sitekey="6LeWx2ogAAAAABZ5B3ZlZbdeJOzIGhmPyo2zsW8d"
                              ></div>
                            </div>
                          </div>
                        </div>
                      </div>
                      <div className="form-group">
                        <button
                          type="submit"
                          className="form-control-btn btn btn-primary submit px-3"
                        >
                          Gửi yêu cầu
                        </button>
                      </div>
                      <div className="login-bottom">
                        <div>
                          <span>
                            Chưa có tài khoản?{' '}
                            <a
                              href="https://kdata.vn/user/register"
                              className="sign-up-now"
                            >
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
                                alt="Zalo"
                                title="Zalo"
                              />
                            </a>
                            <a href="#">
                              <img
                                src="https://kdata.vn/kdata/images/banner/Telegram.png"
                                alt="Telegram"
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
                                alt="Facebook"
                                title="Facebook"
                              />
                            </a>
                          </div>
                          <div className="text-social">
                            <a href="#">
                              Hướng dẫn liên kết tài khoản qua: Zalo, Telegram,
                              Facebook
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
        </div>
      </section>
    </div>
  );
};

export default ForgetPassword;