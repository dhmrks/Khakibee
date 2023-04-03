/* eslint-disable react/jsx-props-no-spreading */
import React from 'react';

import elTr from '../locales/el.json';
import enTr from '../locales/en.json';

const translations = {
  en: enTr,
  el: elTr,
};

const fallbackLng = 'el';

const getLocale = (lng) => {
  const lngExists = Object.prototype.hasOwnProperty.call(translations, lng);

  return {
    lng: lngExists ? lng : fallbackLng,
    translation: lngExists ? translations[lng] : translations[fallbackLng],
  };
};

function withTranslation(WrappedComponent) {
  return class extends React.Component {
    constructor(props) {
      super(props);

      const lng = window.location.pathname.split('/').pop();
      const lc = getLocale(lng);
      this.state = lc;
    }

    setLang = (lng) => {
      const lc = getLocale(lng);
      this.setState(lc);
    };

    render() {
      const { lng, translation } = this.state;
      return <WrappedComponent setLang={this.setLang} lng={lng} tr={translation} {...this.props} />;
    }
  };
}

export default withTranslation;
