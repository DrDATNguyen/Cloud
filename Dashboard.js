import React, { useState } from 'react';
import { Link } from 'react-router-dom';
import RegisterForm from './Login';
import './css/Dashboard.css';

const DashboardUser = () => {
  return (
    <div className="for-home-page dashboard-user">
      <section className="content">
        <header>
          <section className="header-top" id="menuscroll">
            <div className="container">
              <div className="main-header-top">
                <div className="row">
                  <div className="col-md-1">
                    <div className="logo">
                      <a href="https://kdata.vn">
                        <img
                          src="https://kdata.vn/kdata/images/icon/logo-KDATA-vector.svg"
                          alt="KDATA"
                          title="KDATA"
                        />
                      </a>
                    </div>
                  </div>
                  <div className="col-md-3"></div>
                  <div className="col-md-8">
                    <div className="row" style={{ margin: 0 }}>
                      <div className="col-md-3 dpnone">
                        <div className="sale">
                          <img
                            src="https://kdata.vn/kdata/images/icon/icon-hot-khuyen-mai.png"
                            alt="Khuyến mãi"
                          />
                          <a
                            href="https://kdata.vn/promotions"
                            target="_blank"
                            rel="nofollow"
                          >
                            KHUYẾN MÃI
                          </a>
                        </div>
                      </div>
                      <div className="col-md-4 pdnone" style={{ display: 'flex' }}>
                        <div className="signin menu-lv1">
                          <img
                            src="https://kdata.vn/kdata/images/icon/dashboar.svg"
                            width="20px"
                            alt="Dashboard"
                          />
                          <a
                            itemProp="url"
                            href="https://kdata.vn/user/dashboard"
                            rel="nofollow"
                          >
                            Dashboard
                          </a>
                        </div>
                        <div className="signin menu-lv1">
                          <img
                            src="https://kdata.vn/kdata/images/icon/logout.svg"
                            width="20px"
                            alt="Log out"
                          />
                          <Link to="/Login" rel="nofollow">Đăng xuất</Link>
                        </div>
                      </div>
                      <div className="col-md-5">
                        <div className="row" style={{ textAlign: 'center' }}>
                          <div className="col-md-4">
                            <div className="dropdown">
                              <button
                                className="btn dropdown-toggle"
                                type="button"
                                data-toggle="dropdown"
                                style={{ width: '80px', color: '#040404' }}
                              >
                                <img
                                  src="https://kdata.vn/kdata/images/icon/icon-vietnam.png"
                                  alt="Tiếng Việt"
                                />{' '}
                                vi
                                <span className="caret "></span>
                              </button>
                              <ul className="dropdown-menu">
                                <li>
                                  <a href="https://kdata.vn">
                                    <img
                                      src="https://kdata.vn/kdata/images/icon/icon-vietnam.png "
                                      alt="Tiếng Việt"
                                    />{' '}
                                    Vi
                                  </a>
                                </li>
                              </ul>
                            </div>
                          </div>
                          <div
                            className="col-md-4"
                            style={{ borderRight: '2px solid #ff3900' }}
                          >
                            <div className="menu-lv1" style={{ color: '#0042a1' }}>
                              SỐ DƯ <br />
                              <strong>0 đ</strong>
                            </div>
                          </div>
                          <div className="col-md-4">
                            <div className="menu-lv1" style={{ color: '#f86c6b' }}>
                              NỢ <br />
                              <strong>0 đ</strong>
                            </div>
                          </div>
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </section>
        </header>
      </section>
    </div>
  );
};



const Sidebar = () => {
  const [isDropdownOpen, setDropdownOpen] = useState(true); // State để điều khiển hiển thị dropdown

  const toggleDropdown = () => {
    setDropdownOpen(!isDropdownOpen);
  };
  
  return (
    <div className="sidebar">
        <div className="logo">
            <a href="https://kdata.vn">
                <img 
                    src="https://kdata.vn/kdata/images/icon/logo-KDATA-vector.svg" 
                    alt="KDATA" 
                    title="KDATA" 
                    width="160px" 
                    height="40px"
                />
            </a>
        </div>
        <div className="sidebar-wrapper">
            <ul className="nav scroll-container">
                <li>
                    <a className="link" href="https://kdata.vn/user/dashboard">
                        <span>
                            <img 
                                src="https://kdata.vn/kdata/images/icon/main.png" 
                                alt="KDATA" 
                                title="KDATA" 
                            />
                        </span>
                        <p>Tổng quan</p>
                    </a>
                </li>
                <li className="nav-item dropdown active">
                    <span className="ripple rippleEffect" 
                          style={{
                              width: '204.988px', 
                              height: '204.988px', 
                              top: '-90.342px', 
                              left: '86.8774px'
                          }}>
                    </span>
                    <a className="dropdown-btn active">
                        <span>
                            <img 
                                src="https://kdata.vn/kdata/images/icon/cloud.png" 
                                alt="KDATA" 
                                title="KDATA" 
                            />
                        </span>
                        <p>Cloud Hosting</p>
                        <i className="fas fa-chevron-down down" aria-hidden="true"></i>
                    </a>
                    <div className="dropdown-container active" style={{ display: 'block' }}>
                        <a className="nav-link" href="https://kdata.vn/user/hosting">Quản lý Hosting</a>
                        <a className="nav-link active" href="https://kdata.vn/user/create/hosting">Đăng ký Hosting</a>
                    </div>
                </li>
                <li className="slider" style={{ left: '160px' }}></li>
            </ul>
        </div>
    </div>
);
 

};

const MainPanel = () => {
  return (
    <div className="main-panel">
      <div className="content">
        <div className="top-main-dashboard">
          <div className="row">
            <div className="col-lg-12">
              <div className="title-main-db">
                <h3>Khởi tạo HOSTING mới</h3>
                <nav aria-label="breadcrumb">
                  <ol className="breadcrumb">
                    <li className="breadcrumb-item">
                      <a href="https://kdata.vn/user/dashboard">
                        <img
                          src="https://kdata.vn/kdata/images/icon-breadcrumb.png"
                          alt="Star"
                        />{" "}
                        Dashboard
                      </a>
                    </li>
                    <li
                      className="breadcrumb-item active"
                      aria-current="page"
                    >
                      Khởi tạo HOSTING
                    </li>
                  </ol>
                </nav>
              </div>
            </div>
          </div>
        </div>
        <div className="row">
          <div className="col-md-12">
            <div
              className="card"
              style={{ marginTop: '-15px', marginBottom: '0' }}
            >
              <div className="col-lg-12 mes_info"></div>
              <div
                className="panel panel-default panel-accent-asbestos"
                style={{ marginBottom: '0', border: 'none' }}
              >
                <div className="card-body" style={{ padding: '0 35px 0 15px' }}>
                  <div className="main-content-service-dashboard">
                    <div className="row">
                      <div className="col-md-12">
                        <div className="row">
                          <div className="col-md-12">
                            <div className="tab-content">
                              <div className="row">
                                <div className="col-md-12">
                                  <label
                                    className="control-label title-label"
                                    htmlFor="package"
                                  >
                                    LOẠI DỊCH VỤ
                                  </label>
                                  <div className="content-build">
                                    <div className="box-carousel">
                                      <section id="section-type-package">
                                        <div className="container-type-package box-container-package">
                                          <div className="show-some">
                                            <div className="row">
                                              <div className="col-style col-xxl-2 col-xl-3 col-lg-3 col-md-3 block-always-show-package spacing-between-items">
                                                <div className="item">
                                                  <div
                                                    className="box action-select-package box-item active"
                                                    data-type-id="4"
                                                    data-type-name="Wordpress Hosting"
                                                  >
                                                    <div className="arrow-left"></div>
                                                    <i
                                                      className="fa fa-check"
                                                      aria-hidden="true"
                                                    ></i>
                                                    <div className="name-1">
                                                      <img
                                                        src="https://kdata.vn/kdata/images/img-service-5.png"
                                                        alt="img-service"
                                                        style={{
                                                          height: '100px',
                                                          width: 'auto',
                                                          marginBottom: '5px',
                                                        }}
                                                      />
                                                    </div>
                                                    <div className="price-1">
                                                      <b style={{ fontSize: '16px' }}>
                                                        Wordpress Hosting
                                                      </b>
                                                    </div>
                                                  </div>
                                                </div>
                                                <br />
                                              </div>
                                              <div className="col-style col-xxl-2 col-xl-3 col-lg-3 col-md-3 block-always-show-package spacing-between-items">
                                                <div className="item">
                                                  <div
                                                    className="box action-select-package box-item"
                                                    data-type-id="13,14,15"
                                                    data-type-name="Cloud Hosting"
                                                  >
                                                    <div className="arrow-left"></div>
                                                    <i
                                                      className="fa fa-check"
                                                      aria-hidden="true"
                                                    ></i>
                                                    <div className="name-1">
                                                      <img
                                                        src="https://kdata.vn/kdata/images/img-service-1.png"
                                                        alt="img-service"
                                                        style={{
                                                          height: '100px',
                                                          width: 'auto',
                                                          marginBottom: '5px',
                                                        }}
                                                      />
                                                    </div>
                                                    <div className="price-1">
                                                      <b style={{ fontSize: '16px' }}>
                                                        Cloud Hosting
                                                      </b>
                                                    </div>
                                                  </div>
                                                </div>
                                                <br />
                                              </div>
                                            </div>
                                          </div>
                                        </div>
                                        <br />
                                      </section>
                                    </div>
                                  </div>
                                </div>
                              </div>
                              <div className="get_package">
                                <div className="row">
                                  <div className="col-md-12">
                                    <label
                                      className="control-label title-label"
                                      htmlFor="package"
                                    >
                                      GÓI DỊCH VỤ
                                    </label>
                                    <div className="content-build">
                                      <div className="box-carousel">
                                        <section id="section-package">
                                          <div className="container-package box-container-package">
                                            <div className="show-some">
                                              <div className="row">
                                                <div className="col-style col-xxl-2 col-xl-4 col-lg-4 col-md-4 block-always-show-package">
                                                  <div className="item">
                                                    <div
                                                      className="box action-select-package box-item active"
                                                      data-month="1"
                                                      data-id="17"
                                                      data-description="Website Space: 1GB SSDDomain: 1Bandwidth: UnlimitedDatabase: UnlimitedSubdomain: Unlimited"
                                                      data-price_backup="0"
                                                      data-original_monthly="15000"
                                                      data-name="WP 01"
                                                    >
                                                      <div className="arrow-left"></div>
                                                      <i
                                                        className="fa fa-check"
                                                        aria-hidden="true"
                                                      ></i>
                                                      <div className="price-1">
                                                        <b>WP 01</b>
                                                        <br />
                                                      </div>
                                                      <div className="line"></div>
                                                      <div style={{ paddingLeft: '10px' }}>
                                                        Website Space: 1GB SSD
                                                        <br />
                                                        Domain: 1
                                                        <br />
                                                        Bandwidth: Unlimited
                                                        <br />
                                                        Database: Unlimited
                                                        <br />
                                                        Subdomain: Unlimited
                                                      </div>
                                                    </div>
                                                  </div>
                                                  <br />
                                                </div>
                                                <div className="col-style col-xxl-2 col-xl-4 col-lg-4 col-md-4 block-always-show-package">
                                                  <div className="item">
                                                    <div
                                                      className="box action-select-package box-item"
                                                      data-month="1"
                                                      data-id="18"
                                                      data-description="Website Space: 3GB SSDDomain: 2Bandwidth: UnlimitedDatabase: UnlimitedSubdomain: Unlimited"
                                                      data-price_backup="0"
                                                      data-original_monthly="30000"
                                                      data-name="WP 02"
                                                    >
                                                      <div className="arrow-left"></div>
                                                      <i
                                                        className="fa fa-check"
                                                        aria-hidden="true"
                                                      ></i>
                                                      <div className="price-1">
                                                        <b>WP 02</b>
                                                        <br />
                                                      </div>
                                                      <div className="line"></div>
                                                      <div style={{ paddingLeft: '10px' }}>
                                                        Website Space: 3GB SSD
                                                        <br />
                                                        Domain: 2
                                                        <br />
                                                        Bandwidth: Unlimited
                                                        <br />
                                                        Database: Unlimited
                                                        <br />
                                                        Subdomain: Unlimited
                                                      </div>
                                                    </div>
                                                  </div>
                                                  <br />
                                                </div>
                                                <div className="col-style col-xxl-2 col-xl-4 col-lg-4 col-md-4 block-always-show-package">
                                                  <div className="item">
                                                    <div
                                                      className="box action-select-package box-item"
                                                      data-month="1"
                                                      data-id="19"
                                                      data-description="Website Space: 5GB SSDDomain: 4Bandwidth: UnlimitedDatabase: UnlimitedSubdomain: Unlimited"
                                                      data-price_backup="0"
                                                      data-original_monthly="59000"
                                                      data-name="WP 03"
                                                    >
                                                      <div className="arrow-left"></div>
                                                      <i
                                                        className="fa fa-check"
                                                        aria-hidden="true"
                                                      ></i>
                                                      <div className="price-1">
                                                        <b>WP 03</b>
                                                        <br />
                                                      </div>
                                                      <div className="line"></div>
                                                      <div style={{ paddingLeft: '10px' }}>
                                                        Website Space: 5GB SSD
                                                        <br />
                                                        Domain: 4
                                                        <br />
                                                        Bandwidth: Unlimited
                                                        <br />
                                                        Database: Unlimited
                                                        <br />
                                                        Subdomain: Unlimited
                                                      </div>
                                                    </div>
                                                  </div>
                                                  <br />
                                                </div>
                                                <div className="col-style col-xxl-2 col-xl-4 col-lg-4 col-md-4 block-always-show-package">
                                                  <div className="item">
                                                    <div
                                                      className="box action-select-package box-item"
                                                      data-month="1"
                                                      data-id="20"
                                                      data-description="Website Space: 7GB SSDDomain: 6Bandwidth: UnlimitedDatabase: UnlimitedSubdomain: Unlimited"
                                                      data-price_backup="0"
                                                      data-original_monthly="79000"
                                                      data-name="WP 04"
                                                    >
                                                      <div className="arrow-left"></div>
                                                      <i
                                                        className="fa fa-check"
                                                        aria-hidden="true"
                                                      ></i>
                                                      <div className="price-1">
                                                        <b>WP 04</b>
                                                        <br />
                                                      </div>
                                                      <div className="line"></div>
                                                      <div style={{ paddingLeft: '10px' }}>
                                                        Website Space: 7GB SSD
                                                        <br />
                                                        Domain: 6
                                                        <br />
                                                        Bandwidth: Unlimited
                                                        <br />
                                                        Database: Unlimited
                                                        <br />
                                                        Subdomain: Unlimited
                                                      </div>
                                                    </div>
                                                  </div>
                                                  <br />
                                                </div>
                                                <div className="col-style col-xxl-2 col-xl-4 col-lg-4 col-md-4 block-always-show-package">
                                                  <div className="item">
                                                    <div
                                                      className="box action-select-package box-item"
                                                      data-month="1"
                                                      data-id="21"
                                                      data-description="Website Space: 10GB SSDDomain: 8Bandwidth: UnlimitedDatabase: UnlimitedSubdomain: Unlimited"
                                                      data-price_backup="0"
                                                      data-original_monthly="109000"
                                                      data-name="WP 05"
                                                    >
                                                      <div className="arrow-left"></div>
                                                      <i
                                                        className="fa fa-check"
                                                        aria-hidden="true"
                                                      ></i>
                                                      <div className="price-1">
                                                        <b>WP 05</b>
                                                        <br />
                                                      </div>
                                                      <div className="line"></div>
                                                      <div style={{ paddingLeft: '10px' }}>
                                                        Website Space: 10GB SSD
                                                        <br />
                                                        Domain: 8
                                                        <br />
                                                        Bandwidth: Unlimited
                                                        <br />
                                                        Database: Unlimited
                                                        <br />
                                                        Subdomain: Unlimited
                                                      </div>
                                                    </div>
                                                  </div>
                                                  <br />
                                                </div>
                                                <div className="col-style col-xxl-2 col-xl-4 col-lg-4 col-md-4 block-always-show-package">
                                                  <div className="item">
                                                    <div
                                                      className="box action-select-package box-item"
                                                      data-month="1"
                                                      data-id="22"
                                                      data-description="Website Space: 15GB SSDDomain: 14Bandwidth: UnlimitedDatabase: UnlimitedSubdomain: Unlimited"
                                                      data-price_backup="0"
                                                      data-original_monthly="159000"
                                                      data-name="WP 06"
                                                    >
                                                      <div className="arrow-left"></div>
                                                      <i
                                                        className="fa fa-check"
                                                        aria-hidden="true"
                                                      ></i>
                                                      <div className="price-1">
                                                        <b>WP 06</b>
                                                        <br />
                                                      </div>
                                                      <div className="line"></div>
                                                      <div style={{ paddingLeft: '10px' }}>
                                                        Website Space: 15GB SSD
                                                        <br />
                                                        Domain: 14
                                                        <br />
                                                        Bandwidth: Unlimited
                                                        <br />
                                                        Database: Unlimited
                                                        <br />
                                                        Subdomain: Unlimited
                                                      </div>
                                                    </div>
                                                  </div>
                                                  <br />
                                                </div>
                                              </div>
                                            </div>
                                          </div>
                                          <br />
                                        </section>
                                      </div>
                                    </div>
                                  </div>
                                </div>
                              </div>
                            </div>
                          </div>
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

const Aside = () => {
  return (
    <div className="aside">
      {/* Hidden Input */}
      <input type="hidden" id="aside-total-price-qty" value="378000" />

      {/* Sidebar Box */}
      <div className="box-sidebar">
        {/* Order Summary Title */}
        <div className="title-head text-uppercase">Order Summary</div>

        {/* Order Price */}
        <div className="content">
          <div className="box-aside">
            <div className="title">Đơn giá:</div>
            <div className="original_monthly">
              <span className="aside-total-price-before">540,000</span>
            </div>
          </div>

          {/* Quantity */}
          <div className="box-aside" style={{ display: 'none' }}>
            Số lượng <span className="x_qty">1</span>
          </div>
        </div>

        {/* Discount Information */}
        <div className="content discount_info" style={{ display: 'block' }}>
          <div className="box-aside">
            <div className="title">Giảm chu kỳ:</div>
            <div className="price_discount">
              <span className="percent-price_discount">Giảm 30%</span>
              <br />
              <span className="aside-total-price_discount">- 162,000</span>
            </div>
          </div>
        </div>

        {/* Coupon Promotion */}
        <div className="content promotion_info" style={{ display: 'none' }}>
          <div className="box-aside">
            <div className="title">Giảm mã Coupon:</div>
            <span
              className="aside-total-price-qty-promotion"
              style={{ fontWeight: 'bold', color: '#000000' }}
            >
              -65,400
            </span>
          </div>
          <div className="box-aside">
            <small>
              <p>
                Giảm 20% hóa đơn đầu tiên các gói Cloud Hosting chu kỳ 3 &amp; 6 tháng tại KDATA
              </p>
            </small>
          </div>
        </div>

        {/* Points Information */}
        <div className="content points_info" style={{ display: 'none' }}>
          <div className="box-aside">
            <div className="title">Điểm tích luỹ:</div>
            <div className="price_discount" style={{ textAlign: 'right' }}>
              <span className="points_discount">- 10</span>
              <br />
              <span
                className="aside-total-price-qty-promotion"
                style={{ fontWeight: 'bold', color: '#000000' }}
              >
                -65,400
              </span>
            </div>
          </div>
          <div className="box-aside">
            <small>
              <p>1 điểm tương đương 10 vnđ</p>
            </small>
          </div>
        </div>

        {/* VAT Section */}
        <div className="content box-aside">
          <div className="title">VAT:</div>
          <span className="aside-total-price-vat" style={{ fontWeight: 'bold', color: '#000000' }}>
            37,800
          </span>
        </div>

        {/* Total Price */}
        <div className="content" style={{ marginBottom: 0 }}>
          <div className="box-aside">
            <div className="title">Thành tiền:</div>
            <span className="aside-total-price" style={{ fontWeight: 'bold', color: '#000000' }}>
              415,800
            </span>
          </div>
          <div className="box-aside">
            <small>Tương đương</small>
            <small className="x_tuong_duong">10,500/ tháng + VAT</small>
          </div>
        </div>

        {/* Currency Unit */}
        <small>(Đơn vị tính: VNĐ)</small>
      </div>

      {/* Deploy Button */}
      <div className="sum_deploy">
        <div className="deploy_now" style={{ textAlign: 'right' }}>
          <div className="update ml-auto mr-auto">
            <button type="submit" className="btn btn-primary btn-block" id="deploy_now">
              Khởi tạo
            </button>
          </div>
        </div>
      </div>
    </div>
  );
  
};

const BillingCycle = () => {
  return (
    <div className="get_billing_cycle">
      <div className="row">
        <div className="col-md-12">
          <div className="row">
            <div className="col-md-12">
              <label className="control-label title-label" htmlFor="package">CHU KỲ</label>
              <div className="content-build">
                <div className="col-md-12">
                  <div className="tab-content">
                    <div className="row">
                      <div className="panel panel-default panel-accent-asbestos" style={{ marginBottom: '10px' }}>
                        <div className="card-body">
                          <div className="row list-item-service">
                            {/* Lựa chọn chu kỳ 3 tháng */}
                            <div className="col-md-4 col-sm-6">
                              <div className="list__item radio_choice_month">
                                <input type="radio" className="radio-btn" name="choice_month" value="3" data-discount="0" id="3month-opt" disabled />
                                <label htmlFor="3month-opt" className="label">3 Tháng</label>
                              </div>
                            </div>

                            {/* Lựa chọn chu kỳ 6 tháng */}
                            <div className="col-md-4 col-sm-6">
                              <div className="list__item radio_choice_month">
                                <input type="radio" className="radio-btn" name="choice_month" value="6" data-discount="0" id="6month-opt" disabled />
                                <label htmlFor="6month-opt" className="label">6 Tháng</label>
                              </div>
                            </div>

                            {/* Lựa chọn chu kỳ 12 tháng */}
                            <div className="col-md-4 col-sm-6">
                              <div className="list__item radio_choice_month">
                                <input type="radio" className="radio-btn" name="choice_month" value="12" data-discount="10" id="12month-opt" />
                                <label htmlFor="12month-opt" className="label">
                                  12 Tháng
                                  <span className="cloud-vps-discount">-10 %</span>
                                </label>
                              </div>
                            </div>

                            {/* Lựa chọn chu kỳ 24 tháng */}
                            <div className="col-md-4 col-sm-6">
                              <div className="list__item radio_choice_month">
                                <input type="radio" className="radio-btn" name="choice_month" value="24" data-discount="15" id="24month-opt" />
                                <label htmlFor="24month-opt" className="label">
                                  24 Tháng
                                  <span className="cloud-vps-discount">-15 %</span>
                                </label>
                              </div>
                            </div>

                            {/* Lựa chọn chu kỳ 36 tháng */}
                            <div className="col-md-4 col-sm-6">
                              <div className="list__item radio_choice_month">
                                <input type="radio" className="radio-btn" name="choice_month" value="36" data-discount="30" id="36month-opt" defaultChecked />
                                <label htmlFor="36month-opt" className="label">
                                  36 Tháng
                                  <span className="cloud-vps-discount">-30%</span>
                                </label>
                              </div>
                            </div>
                          </div>
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

const Dashboard = () => {
  return (
    <div className="dashboard-container">
      <DashboardUser />
      <div className="content">
        <Sidebar />
       <MainPanel/>
       <Aside/>
       <BillingCycle />
      </div>
    </div>
  );
};

export default Dashboard;