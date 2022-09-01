import './Home.scss';

import useSEO from '../../hooks/useSEO';
import meli_mobile_logo from '../../assets/meli_mobile_logo.png';

export default function HomePage() {    
    const ogImage = meli_mobile_logo

    useSEO({
        title: 'Mercado Libre - Donde comprar y vender de todo',
        description: 'Comprá productos con Envío Gratis en el día en Mercado Libre Uruguay. Encontrá miles de marcas y productos a precios increíbles.',
        ogImage: ogImage
    })

    return (
        <main>

        </main>
    );
}
