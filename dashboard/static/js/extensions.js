htmx.defineExtension('echarts', {
    transformResponse: function (text, xhr, elt) {
        // parse json data
        var data = JSON.parse(text);

        // fetch echart element
        var option = data;
        chartsMap.get(data.id).chart.setOption(option);

        return "";
    }
});

/*
document.addEventListener('DOMContentLoaded', () => {
    const dropdowns = document.querySelectorAll('.dropdown');

    dropdowns.forEach(dropdown => {
        const btn = dropdown.querySelector('.dropdown-btn');
        const menu = dropdown.querySelector('.dropdown-menu');

        btn.addEventListener('click', () => {
            // Close all other dropdowns
            dropdowns.forEach(otherDropdown => {
                if (otherDropdown !== dropdown) {
                    otherDropdown.querySelector('.dropdown-menu').classList.add('hidden');
                }
            });

            // Toggle this dropdown
            menu.classList.toggle('hidden');
        });
    });

    // Close all dropdowns when clicking outside
    window.addEventListener('click', (event) => {
        if (!event.target.matches('.dropdown-btn')) {
            dropdowns.forEach(dropdown => {
                const menu = dropdown.querySelector('.dropdown-menu');
                if (!menu.contains(event.target)) {
                    menu.classList.add('hidden');
                }
            });
        }
    });
});*/
