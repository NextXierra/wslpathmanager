import * as App from './wailsjs/go/main/App.js';

document.addEventListener('DOMContentLoaded', async () => {
    const distroSelect = document.getElementById('distroSelect');
    const scanBtn = document.getElementById('scanBtn');
    const loading = document.getElementById('loading');
    const results = document.getElementById('results');
    const emptyState = document.getElementById('emptyState');
    const toolsList = document.getElementById('toolsList');
    const selectAllBtn = document.getElementById('selectAllBtn');
    const unselectAllBtn = document.getElementById('unselectAllBtn');
    const saveBtn = document.getElementById('saveBtn');

    let currentTools = [];

    // Load Distros
    try {
        const distros = await App.GetDistros();
        if (distros && distros.length > 0) {
            distroSelect.innerHTML = '';
            distros.forEach(distro => {
                const option = document.createElement('option');
                option.value = distro.name;
                option.textContent = distro.name;
                distroSelect.appendChild(option);
            });
            // trigger scan on load like original app
            scanPaths();
        }
    } catch (err) {
        console.error("Failed to load distros", err);
    }

    distroSelect.addEventListener('change', () => {
        scanPaths();
    });

    scanBtn.addEventListener('click', scanPaths);

    async function scanPaths() {
        const distroName = distroSelect.value;
        if (!distroName) return;

        // Show loading state
        emptyState.style.display = 'none';
        results.style.display = 'none';
        saveBtn.style.display = 'none';
        loading.style.display = 'inline-block';
        
        try {
            const paths = await App.ScanPaths(distroName);
            currentTools = paths || [];
            
            renderTools();
            
            loading.style.display = 'none';
            if (currentTools.length > 0) {
                results.style.display = 'flex';
                saveBtn.style.display = 'flex';
            } else {
                emptyState.style.display = 'block';
            }
        } catch (err) {
            console.error(err);
            loading.style.display = 'none';
            emptyState.style.display = 'block';
        }
    }

    function renderTools() {
        toolsList.innerHTML = '';
        currentTools.forEach((tool, index) => {
            const item = document.createElement('div');
            item.className = 'tool-item';

            const label = document.createElement('label');
            label.className = 'tool-label';

            const checkbox = document.createElement('input');
            checkbox.type = 'checkbox';
            checkbox.checked = tool.isSelected;
            checkbox.dataset.index = index;
            
            checkbox.addEventListener('change', (e) => {
                currentTools[e.target.dataset.index].isSelected = e.target.checked;
            });

            const span = document.createElement('span');
            span.textContent = tool.toolName;

            label.appendChild(checkbox);
            label.appendChild(span);

            const pathDiv = document.createElement('div');
            pathDiv.className = 'tool-path';
            pathDiv.textContent = tool.wslPath;

            item.appendChild(label);
            item.appendChild(pathDiv);
            
            toolsList.appendChild(item);
        });
    }

    selectAllBtn.addEventListener('click', () => {
        currentTools.forEach(t => t.isSelected = true);
        renderTools();
    });

    unselectAllBtn.addEventListener('click', () => {
        currentTools.forEach(t => t.isSelected = false);
        renderTools();
    });

    saveBtn.addEventListener('click', async () => {
        const distroName = distroSelect.value;
        if (!distroName) return;

        saveBtn.disabled = true;
        saveBtn.textContent = 'Saving...';

        const selectedTools = currentTools
            .filter(t => t.isSelected)
            .map(t => t.toolName);

        try {
            await App.SaveSettings(distroName, selectedTools);
            saveBtn.textContent = 'Saved!';
            setTimeout(() => {
                saveBtn.textContent = 'Save & Apply Settings';
                saveBtn.disabled = false;
            }, 2000);
        } catch (err) {
            console.error(err);
            saveBtn.textContent = 'Error saving';
            setTimeout(() => {
                saveBtn.textContent = 'Save & Apply Settings';
                saveBtn.disabled = false;
            }, 2000);
        }
    });

    const settingsBtn = document.getElementById('settingsBtn');
    const settingsModal = document.getElementById('settingsModal');
    const closeSettingsBtn = document.getElementById('closeSettingsBtn');
    const themeSelect = document.getElementById('themeSelect');

    settingsBtn.addEventListener('click', () => {
        settingsModal.style.display = 'flex';
    });

    closeSettingsBtn.addEventListener('click', () => {
        settingsModal.style.display = 'none';
    });

    themeSelect.addEventListener('change', (e) => {
        if (e.target.value === 'light') {
            document.documentElement.classList.add('light');
            document.body.classList.add('light');
            document.body.classList.remove('dark');
        } else {
            document.documentElement.classList.remove('light');
            document.body.classList.remove('light');
            document.body.classList.add('dark');
        }
        localStorage.setItem('theme', e.target.value);
    });

    const savedTheme = localStorage.getItem('theme') || 'dark';
    themeSelect.value = savedTheme;
    if (savedTheme === 'light') {
        document.documentElement.classList.add('light');
        document.body.classList.add('light');
        document.body.classList.remove('dark');
    }
});
